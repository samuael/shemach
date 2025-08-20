package rest

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/shemach/backend/pkg/constants/types"
	"github.com/samuael/shemach/backend/pkg/resource"
	"github.com/samuael/shemach/backend/platforms/helper"
)

type IResourceHandler interface {
	UploadPropertyImage(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetImageByID(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	// GetBlurredImage(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type ResourceHandler struct {
	Service resource.IResourceService
}

func NewResourceHandler(service resource.IResourceService) IResourceHandler {
	return &ResourceHandler{
		Service: service,
	}
}

// UploadPropertyImage : authenticated endpoint and user session is required
func (rHandler *ResourceHandler) UploadPropertyImage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseMultipartForm(99999999999)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "bad input: " + err.Error(), Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	propertyDetailIDs := r.MultipartForm.Value["property_detail_id"]
	if len(propertyDetailIDs) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "bad input: property detail id is required", Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	propertyDetailID, err := strconv.ParseUint(propertyDetailIDs[0], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: fmt.Sprintf("bad input: invalid property detail id %q", propertyDetailIDs[0]), Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	session, err := SessionFromContext(ctx)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(helper.MarshalThis(&types.StatusMsg{Status: types.ScUnauthorized, Error: "unauthorized"}))
		return
	}
	var isCoverPicture bool
	isCoverString := r.MultipartForm.Value["isCover"]
	if len(isCoverString) == 1 || isCoverString[0] == "" {
		isCoverPicture, err = strconv.ParseBool(isCoverString[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			messageBody := helper.MarshalThis(&types.StatusMsg{Error: "bad input: " + err.Error(), Status: types.ScBadRequest})
			w.Write(messageBody)
			return
		}
	}
	status, err := rHandler.Service.CheckPropertyImageUploadEligibility(ctx, propertyDetailID, session.ID, isCoverPicture, 5)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "internal error: " + err.Error(), Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	switch status {
	case 0:
	case 1:
		w.WriteHeader(http.StatusNotFound)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "property detail not found", Status: types.ScNotFound})
		w.Write(messageBody)
		return
	case 2:
		panic("http.StatusUnauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "you are not authorized", Status: types.ScUnauthorized})
		w.Write(messageBody)
		return
	case 3:
		panic("http.StatusForbidden")
		w.WriteHeader(http.StatusForbidden)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "property images exceeded maximum quantity of 5", Status: types.ScFailed})
		w.Write(messageBody)
		return
	default:
		panic("http.StatusConflict")
		w.WriteHeader(http.StatusConflict)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "property cover picture already uploaded: " + strconv.FormatUint(status, 10), Status: types.ScConflict})
		w.Write(messageBody)
		return
	}

	// What do I do here?

	image, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusBadRequest)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "valid image 'file' required", Status: types.ScBadRequest})
		w.Write(messageBody)
		return
	}
	defer image.Close()
	if !helper.IsImage(header.Filename) {
		fileSplit := strings.Split(header.Filename, ".")
		extension := header.Filename
		if len(fileSplit) > 1 && len(fileSplit[0]) > 0 {
			extension = fileSplit[0]
		}
		panic("StatusUnsupportedMediaType")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "unsupported media type " + extension, Status: types.ScUnsupportedMediaType})
		w.Write(messageBody)
		return
	}
	newName := "images/property/" + helper.GenerateRandomString(5, helper.CHARACTERS) + "." + helper.GetExtension(header.Filename)
	var newImage *os.File
	newImage, err = os.Create(os.Getenv("ASSETS_DIRECTORY") + newName)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "internal server error please try again" + err.Error(), Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	defer newImage.Close()

	_, err = io.Copy(newImage, image)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "internal server error please try again" + err.Error(), Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}

	imageID, err := rHandler.Service.SavePropertyImage(ctx, newName, propertyDetailID, session.ID, isCoverPicture, 0)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "internal server error saving the image; please try again" + err.Error(), Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	// // SavePropertyImage saves an image information into the property detail
	// -- 1 if the property not found
	// -- 2 owner of the property mismatch mismatch
	// -- 3 insertion error
	// -- 4 cover picture update error
	// -- 5 pictures update error
	// -- bigint successful insertion.
	switch imageID {
	case 1, 2:
		// already handled above
	case 3, 4, 5:
		panic("internal server error saving the image; please try again later")
		w.WriteHeader(http.StatusInternalServerError)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "internal server error saving the image; please try again later", Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	// Upload Succesful
	w.WriteHeader(http.StatusOK)
	w.Write(helper.MarshalThis(&types.StatusMsg{
		Status:  types.ScOK,
		Data:    map[string]uint64{"id": imageID},
		Success: true,
	}))
}

func (rHandler *ResourceHandler) GetImageByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	imgID, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO: Role check and other filters will be implemented here.
	path, err := rHandler.Service.GetImagePath(r.Context(), imgID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	file, err := os.Open(os.Getenv("ASSETS_DIRECTORY") + path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set headers
	w.Header().Set("Content-Disposition", "attachment; filename=\""+path+"\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
