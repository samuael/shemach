package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/agent"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/crop"
	"github.com/samuael/agri-net/agri-net-backend/pkg/merchant"
	"github.com/samuael/agri-net/agri-net-backend/pkg/product"
	"github.com/samuael/agri-net/agri-net-backend/pkg/resource"
	"github.com/samuael/agri-net/agri-net-backend/pkg/store"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/resource_handler"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type ICropHandler interface {
	CreateProduct(c *gin.Context)
	UploadProductImages(c *gin.Context)
	Getposts(c *gin.Context)
	GetPostByID(c *gin.Context)
	GetMyPosts(c *gin.Context)
}
type CropHandler struct {
	Service         crop.ICropService
	ProductService  product.IProductService
	StoreService    store.IStoreService
	MerchantService merchant.IMerchantService
	AgentService    agent.IAgentService
	ResourceService resource.IResourceService
}

func NewCropHandler(
	service crop.ICropService,
	productservice product.IProductService,
	storeService store.IStoreService,
	merchantService merchant.IMerchantService,
	agentService agent.IAgentService,
	resourceservice resource.IResourceService,
) ICropHandler {
	return &CropHandler{
		Service:         service,
		ProductService:  productservice,
		StoreService:    storeService,
		MerchantService: merchantService,
		AgentService:    agentService,
		ResourceService: resourceservice,
	}
}

func (chandler *CropHandler) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	if !(session.Role == state.MERCHANT || session.Role == state.AGENT) {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}
	input := &struct {
		TypeID            uint    `json:"type_id"`
		SellingPrice      float64 `json:"selling_price"`
		StoreID           uint64  `json:"store_id"`
		Description       string  `json:"description"`
		RemainingQuantity uint64  `json:"quantity"`
		Negotiable        bool    `json:"negotiable_price"`
	}{}
	res := &struct {
		StatusCode int               `json:"status_code"`
		Msg        string            `json:"msg"`
		Crop       *model.Crop       `json:"crop"`
		Errors     map[string]string `json:"errors,omitempty"`
	}{
		Errors: map[string]string{},
	}
	jdecode := json.NewDecoder(c.Request.Body)
	er := jdecode.Decode(input)
	if er != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "bad request input")
		c.JSON(res.StatusCode, res)
		return
	}
	failed := false
	ack, desc := chandler.ProductService.GetProductInfoByID(int(input.TypeID))
	if ack == "" || desc == "" {
		res.Errors["type"] = translation.Translate(session.Lang, "product type with this id does not exist")
		failed = true
	}
	if input.RemainingQuantity <= 0 {
		res.Errors["remaining_quantity"] = translation.Translate(session.Lang, "unacceptable remaining quantity")
		failed = true
	}
	if input.SellingPrice <= 0 {
		res.Errors["selling_price"] = translation.Translate(session.Lang, "invalid selling price")
		failed = true
	}
	var address *model.Address
	isstoreOwned := true
	if session.Role == state.MERCHANT && input.StoreID <= 0 {
		res.Errors["store_id"] = translation.Translate(session.Lang, "the store id is missing ")
		failed = true
	} else if session.Role == state.MERCHANT {
		store, er := chandler.StoreService.GetStoreByID(ctx, input.StoreID)
		if store.OwnerID != session.ID || er != nil {
			res.StatusCode = http.StatusBadRequest
			res.Errors["store"] = translation.Translate(session.Lang, "you are not authorized to access this store")
			failed = true
		}
		address = store.Address
	} else if session.Role == state.AGENT {
		isstoreOwned = false
		ad, er := chandler.AgentService.GetAgentsAddress(ctx, int(session.ID))
		if er != nil {
			res.StatusCode = http.StatusNotImplemented
			res.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
			c.JSON(res.StatusCode, res)
			return
		} else {
			address = ad
		}
	}
	if address == nil {
		res.StatusCode = http.StatusNotImplemented
		res.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(res.StatusCode, res)
		return
	}
	if failed {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "bad request payload")
		c.JSON(res.StatusCode, res)
		return
	}
	crop := &model.Crop{
		TypeID:            input.TypeID,
		Description:       input.Description,
		Address:           address,
		RemainingQuantity: input.RemainingQuantity,
		Negotiable:        input.Negotiable,
		StoreID: func() uint64 {
			if isstoreOwned {
				return input.StoreID
			}
			return 0
		}(),
		StoreOwned:   isstoreOwned,
		SellingPrice: input.SellingPrice,
		AddressRef:   uint64(address.ID),
		AgentID: func() uint64 {
			if session.Role == state.AGENT {
				return session.ID
			}
			return 0
		}(),
	}
	status, err := chandler.Service.CreateCrop(ctx, crop)
	if err != nil || status < 0 {
		if status == -1 {
			res.StatusCode = http.StatusNotFound
			res.Msg = translation.Translate(session.Lang, "missing address information")
		} else if status == -2 {
			res.StatusCode = http.StatusUnauthorized
			res.Msg = translation.Translate(session.Lang, "you are not authorized")
		} else {
			res.Msg = translation.Translate(session.Lang, "internal server error")
			res.StatusCode = http.StatusInternalServerError
		}
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusCreated
	res.Msg = translation.Translate(session.Lang, "created")
	res.Crop = crop
	res.Errors = nil
	c.JSON(res.StatusCode, res)
}

type MultipartData struct {
	File   multipart.File
	Header *multipart.FileHeader
	Error  error
}

// UploadProductImages ...
func (chandler CropHandler) UploadProductImages(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Minute*1))
	session := ctx.Value("session").(*model.Session)
	res := &struct {
		Msg               string   `json:"msg"`
		StatusCode        int      `json:"status_code"`
		BluredImageRoutes []string `json:"blured_image_routes"`
		ImageRoutes       []string `json:"image_routes"`
		Error             string   `json:"error,omitempty"`
	}{}
	err := c.Request.ParseMultipartForm(9999999999999999)
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt(err.Error()) //"  bad multipart for file file Size Exceeds the limit")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	postID, err := strconv.Atoi(c.Param("postid"))
	if err != nil || postID <= 0 {
		res.Error = translation.Translate(session.Lang, "bad query , missing post id")
		res.StatusCode = http.StatusBadRequest
		c.JSON(res.StatusCode, res)
		return
	}
	post, er := chandler.Service.GetPostByID(ctx, uint64(postID))
	if er != nil {
		res.Error = translation.Translate(session.Lang, "crop post information with this id doesn't exist")
		res.StatusCode = http.StatusNotFound
		c.JSON(res.StatusCode, res)
		return
	}
	var postImageFiles = map[int]*os.File{}
	postMultipartFiles := map[int]*MultipartData{}
	filenames := map[int]string{}
	blurredFilenames := map[int]string{}
	for a := 1; a <= 5; a++ {
		sf := &MultipartData{}
		sf.File, sf.Header, sf.Error = c.Request.FormFile("image" + strconv.Itoa(a))
		if sf.File == nil || sf.Header == nil || sf.Error != nil {
			if sf.Error != nil {
				log.Println(sf.Error.Error())
			}
			log.Println("Something has happened")
			continue
		}
		defer sf.File.Close()
		if !helper.IsImage(sf.Header.Filename) {
			res.Error = translation.Translate(session.Lang, "only image files are allowed")
			res.StatusCode = http.StatusUnsupportedMediaType
			c.JSON(res.StatusCode, res)
			return
		}
		filename :=
			state.POST_IMAGES_RELATIVE_PATH +
				helper.GenerateRandomString(6, helper.CHARACTERS) +
				"." + helper.GetExtension(sf.Header.Filename)
		file, ee := os.Create(os.Getenv("ASSETS_DIRECTORY") + filename)
		if ee != nil {
			println(ee.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Error = " internal server error "
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		_, era := io.Copy(file, sf.File)
		if era != nil {
			log.Println(era.Error())
			continue
		}
		postImageFiles[a] = file
		defer postImageFiles[a].Close()
		postMultipartFiles[a] = sf
		filenames[a] = filename
	}
	if len(postImageFiles) == 0 {
		res.Error = translation.TranslateIt("no image file was uploaded")
		res.StatusCode = http.StatusBadRequest
		c.JSON(res.StatusCode, res)
		return
	}
	for key, saf := range filenames {
		url, er := resource_handler.GetBlurredImage(saf)
		if er != nil {
			os.Remove(os.Getenv("ASSETS_DIRECTORY") + filenames[key])
			delete(filenames, key)
			delete(postImageFiles, key)
			delete(postMultipartFiles, key)
		}
		blurredFilenames[key] = url
	}
	var images []*model.PostImg
	for a := range filenames {
		img := &model.PostImg{
			Resource:       filenames[a],
			OwnerID:        int(session.ID),
			BlurredRe:      blurredFilenames[a],
			Authorized:     false,
			Authorizations: 31,
			OwnerRole: func() int {
				switch session.Role {
				case state.AGENT:
					return state.AGENT_ROLE_INT
				case state.MERCHANT:
					return state.MERCHANT_ROLE_INT
				case state.ADMIN:
					return state.ADMIN_ROLE_INT
				case state.INFO_ADMIN:
					return state.INFOADMIN_ROLE_INT
				case state.SUPERADMIN:
					return state.SUPERADMIN_ROLE_INT
				default:
					return state.ROLE_ALL
				}
			}(),
		}
		images = append(images, img)
	}
	err = chandler.ResourceService.SaveImagesResources(ctx, images)
	if err != nil {
		res.Error = translation.Translate(session.Lang, err.Error()+"internal problem, please try again later!")
		res.StatusCode = http.StatusInternalServerError
		c.JSON(res.StatusCode, res)
		return
	}
	for i := range images {
		if images[i].ID != 0 {
			post.Images = append(post.Images, images[i].ID)
		}
	}
	err = chandler.Service.SaveNewPostImages(ctx, post.ID, post.Images)
	if err != nil {
		res.Error = translation.Translate(session.Lang, err.Error()+"internal problem, please try again later!")
		res.StatusCode = http.StatusInternalServerError
		c.JSON(res.StatusCode, res)
		return
	}
	imageRoutes := []string{}
	blurredImageRoutes := []string{}
	for c := range images {
		imageRoutes = append(imageRoutes, "post/image/"+strconv.Itoa(images[c].ID))
		blurredImageRoutes = append(blurredImageRoutes, "post/image/"+strconv.Itoa(images[c].ID)+"/blurred/")
	}
	res.StatusCode = http.StatusOK
	res.Msg = fmt.Sprintf("%d %s", len(images), translation.Translate(session.Lang, "images uploaded successfuly"))
	res.BluredImageRoutes = blurredImageRoutes
	res.ImageRoutes = imageRoutes
	c.JSON(res.StatusCode, res)
}

func (chandler *CropHandler) GetMyPosts(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	res := &struct {
		Msg        string        `json:"msg"`
		StatusCode int           `json:"status_code"`
		Posts      []*model.Crop `json:"posts"`
	}{}
	offset, er := strconv.Atoi(c.Query("offset"))
	if er != nil {
		offset = 0
	}
	limit, er := strconv.Atoi(c.Query("limit"))
	if er != nil {
		limit = offset + 10
	}
	var posts []*model.Crop
	var stores []*model.Store

	if session.Role == state.MERCHANT {
		stores, er = chandler.StoreService.GetMerchantStores(ctx, session.ID)
		if er != nil {
			res.StatusCode = http.StatusNotFound
			res.Msg = translation.TranslateIt("posts not found")
			c.JSON(res.StatusCode, res)
			return
		}
		storeIDS := []uint64{}
		for x := range stores {
			storeIDS = append(storeIDS, stores[x].ID)
		}
		posts, er = chandler.Service.GetMerchantPosts(ctx, storeIDS, uint(offset), uint(limit))
		if er != nil {
			println("The Error ", er.Error())
			if strings.Contains(er.Error(), "too many connection") {
				res.StatusCode = http.StatusInternalServerError
				res.Msg = translation.Translate(session.Lang, "internal connection problem")
			} else {
				res.StatusCode = http.StatusNotFound
				res.Msg = translation.Translate(session.Lang, "no products found")
			}
			c.JSON(res.StatusCode, res)
			return
		}
	} else {
		posts, er = chandler.Service.GetAgentPosts(ctx, session.ID, uint(offset), uint(limit))
	}
	if er != nil {
		println(er.Error())
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "posts not found")
		c.JSON(res.StatusCode, res)
		return
	}
	res.Posts = posts
	res.StatusCode = http.StatusOK
	res.Msg = translation.TranslateIt("fetched")
	c.JSON(res.StatusCode, res)
}

// Getposts   returns the list of  crops posted in the system.
func (chandler *CropHandler) Getposts(c *gin.Context) {
	ctx := c.Request.Context()
	offset, er := strconv.Atoi(c.Query("offset"))
	res := &struct {
		Msg        string        `json:"msg,omitempty"`
		StatusCode int           `json:"status_code"`
		Posts      []*model.Crop `json:"posts,omitempty"`
	}{}
	if er != nil {
		offset = 0
	}
	limit, er := strconv.Atoi(c.Query("limit"))
	if er != nil {
		limit = offset + 10
	}

	posts, er := chandler.Service.GetPosts(ctx, uint(offset), uint(limit))
	if er != nil {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.TranslateIt("posts not found")
	}
	res.Posts = posts
	res.StatusCode = http.StatusOK
	res.Msg = translation.TranslateIt("fetched")
	c.JSON(res.StatusCode, res)
}

// GetPostByID(c *gin.Context)
func (chandler *CropHandler) GetPostByID(c *gin.Context) {
	ctx := c.Request.Context()
	res := &struct {
		StatusCode int         `json:"status_code"`
		Msg        string      `json:"msg"`
		Post       *model.Crop `json:"post"`
	}{}
	session := ctx.Value("session").(*model.Session)
	id, er := strconv.Atoi(c.Param("id"))
	if er != nil || id <= 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "bad request payload")
		c.JSON(res.StatusCode, res)
		return
	}
	post, er := chandler.Service.GetPostByID(ctx, uint64(id))
	if er != nil {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "post not found")
		c.JSON(res.StatusCode, res)
		return
	}
	res.Post = post
	res.StatusCode = http.StatusOK
	res.Msg = translation.Translate(session.Lang, "post found")
	c.JSON(res.StatusCode, res)
}
