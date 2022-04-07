package rest

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/category"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/platforms"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
)

type ICategoryHandler interface {
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	UpdateCategoryPicture(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	GetCategories(c *gin.Context)
	SetCategoryFee(c *gin.Context)
	GetCategoryStudents(c *gin.Context)
}

type CategoryHandler struct {
	CategorySer category.ICategoryService
}

func NewCategoryHandler(catser category.ICategoryService) ICategoryHandler {
	return &CategoryHandler{
		CategorySer: catser,
	}
}

// CreateCategory
func (cathandler *CategoryHandler) CreateCategory(c *gin.Context) {
	input := &struct {
		Title      string `json:"title"`
		ShortTitle string `json:"short_title"`
	}{}
	res := &struct {
		Msg      string          `json:"msg"`
		Category *model.Category `json:"category"`
	}{
		Category: &model.Category{},
	}
	if ers := c.BindJSON(input); ers != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res.Category.Title = input.Title
	res.Category.ShortTitle = input.ShortTitle
	res.Category.RoundsCount = 0
	ctx := c.Request.Context()
	res.Category.CreatedAt = platforms.GetCurrentEthiopianTime()
	ctx = context.WithValue(ctx, "category", res.Category)
	// check whether a category having simmilar title and short_title exist or not.
	// here the method is get category by title and short_title.
	if cat := cathandler.CategorySer.GetCategoryByContent(ctx); cat != nil {
		res.Msg = "category with this title already exist!"
		res.Category = cat
		c.JSON(http.StatusConflict, res)
		return
	}
	if cat, err, status := cathandler.CategorySer.Create(ctx); cat != nil && err == nil && status == state.DT_STATUS_OK {
		res.Category = cat
		res.Msg = "Category created succesfuly!"
		c.JSON(http.StatusOK, res)
		return
	} else if status == state.DT_STATUS_RECORD_ALREADY_EXIST {
		res.Msg = "category with simmilar data already exist"
		c.JSON(http.StatusConflict, res)
		return
	}
	res.Msg = "category creation was not succesful"
	c.JSON(http.StatusInternalServerError, res)
}

// UpdateCategory ...
func (cathandler *CategoryHandler) UpdateCategory(c *gin.Context) {
	input := &struct {
		ID         uint64 `json:"id"`
		Title      string `json:"title,omitempty"`
		ShortTitle string `json:"short_title,omitempty"`
	}{}
	resp := &struct {
		Msg string `json:"msg"`
	}{"bad input value !\n the Json Has to include id , title  , and short_title"}
	if er := c.BindJSON(input); er != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "category_id", input.ID)
	res, er, code := cathandler.CategorySer.GetCategoryByID(ctx)
	if er != nil || res == nil || code != state.DT_STATUS_OK {
		resp.Msg = "category with this id doesn't exist!"
		c.JSON(http.StatusNotFound, resp)
		return
	}
	res.ID = input.ID
	changed := false
	if input.Title != res.Title {
		if input.Title != "" && input.Title != res.Title {
			res.Title = input.Title
			changed = true
		}
	}
	if input.ShortTitle != res.ShortTitle {
		if input.ShortTitle != "" && input.ShortTitle != res.ShortTitle {
			res.ShortTitle = input.ShortTitle
			changed = true
		}
	}
	if !changed {
		// Not Modified header doesn't have a response body
		resp.Msg = "no change is made."
		c.JSON(http.StatusNotModified, resp)
		return
	}
	stcode := 0
	res.CreatedAt = platforms.GetCurrentEthiopianTime()
	ctx = context.WithValue(ctx, "category", res)
	if err, code := cathandler.CategorySer.UpdateCategory(ctx); err == nil && code == state.DT_STATUS_OK {
		c.JSON(http.StatusOK, res)
		return
	} else {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "category_title_key" (SQLSTATE 23505)` {
			resp.Msg = "category with this title already exist"
			stcode = http.StatusConflict
		} else {
			resp.Msg = "category update was not successful"
			stcode = http.StatusInternalServerError
		}
	}
	c.JSON(stcode, resp)
}

// UpdateCategoryPicture   ... creates a category
func (cathandler *CategoryHandler) UpdateCategoryPicture(c *gin.Context) {
	var header *multipart.FileHeader
	var erro error
	// var oldImage string
	eres := &struct {
		Err string `json:"err"`
	}{
		Err: "bad request",
	}
	categoryID, era := strconv.Atoi(c.Query("id"))
	if era != nil || categoryID <= 0 {
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	erro = c.Request.ParseMultipartForm(99999999999)
	if erro != nil {
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	image, header, erro := c.Request.FormFile("image")
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	defer image.Close()
	ctx := c.Request.Context()
	if helper.IsImage(header.Filename) {
		newName := "images/category/" + helper.GenerateRandomString(5, helper.CHARACTERS) + "." + helper.GetExtension(header.Filename)
		var newImage *os.File
		if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
			newImage, erro = os.Create(os.Getenv("ASSETS_DIRECTORY") + newName)
		} else {
			newImage, erro = os.Create(os.Getenv("ASSETS_DIRECTORY") + "/" + newName)
		}
		if erro != nil {
			println(" 177", erro.Error())
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		defer newImage.Close()
		ctx = context.WithValue(ctx, "category_id", categoryID)
		oldImage, err := cathandler.CategorySer.GetImageUrl(ctx)
		if err != nil {
			println("185: ", err.Error())
			c.JSON(http.StatusInternalServerError, eres)
			return
		}
		_, er := io.Copy(newImage, image)
		if er != nil {
			println(" 191: ", er.Error())
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		ncon := context.WithValue(c.Request.Context(), "category_id", categoryID)
		ncon = context.WithValue(ncon, "image_url", newName)
		success := cathandler.CategorySer.ChangeCategoryImageUrl(ncon)
		if success {
			if oldImage != "" {
				if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
					er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + oldImage)
				} else {
					er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + "/" + oldImage)
				}
			}
			c.JSON(http.StatusOK, &model.ShortSuccess{Msg: newName})
			return
		}
		if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
			er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + newName)
		} else {
			er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + "/" + newName)
		}
		println("214: ", "Bro Updating was not succesful")
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		println("216: ")
		c.Writer.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

// DeleteCategory ...
func (cathandler *CategoryHandler) DeleteCategory(c *gin.Context) {
	res := &struct {
		Msg string `json:"msg"`
	}{
		Msg: "bad request value param 'id' should be >= 1 ",
	}
	catid, er := strconv.Atoi(c.Request.FormValue("id"))
	if catid <= 0 || er != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "category_id", uint64(catid))
	if success := cathandler.CategorySer.DeleteCategoryByID(ctx); success == nil {
		res.Msg = "Category with id " + strconv.Itoa(int(catid)) + " deleted succesfuly"
		c.JSON(http.StatusOK, res)
		return
	} else {
		res.Msg = "category with this id doesn't exist!"
		c.JSON(http.StatusNotFound, res)
		return
	}
}

// GetCategoryByID ...
func (cathandler *CategoryHandler) GetCategoryByID(c *gin.Context) {

	res := &struct {
		Msg      string          `json:"msg"`
		Category *model.Category `json:"category"`
	}{
		Msg: "bad request",
	}
	id, er := strconv.Atoi(c.Request.FormValue("id"))
	log.Println(id)
	if er != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "category_id", uint64(id))
	if category, er, code := cathandler.CategorySer.GetCategoryByID(ctx); er == nil && category != nil && code == state.DT_STATUS_OK {
		res.Msg = "ok"
		res.Category = category
		c.JSON(http.StatusOK, res)
		return
	}
	res.Msg = "category with this id not found"
	c.JSON(http.StatusNotFound, res)
}

// GetCategories ...
func (cathandler *CategoryHandler) GetCategories(c *gin.Context) {
	if categories, er := cathandler.CategorySer.GetCategories(c.Request.Context()); er == nil {
		c.JSON(http.StatusOK, categories)
		return
	}
	c.JSON(http.StatusNotFound, []interface{}{})
}

func (cathandler *CategoryHandler) SetCategoryFee(c *gin.Context) {
	// In this method I am gonna set the training fee for the category.
	in := &struct {
		CategoryID uint64  `json:"id"`
		Amount     float64 `json:"amount"`
	}{}
	eres := &model.ErMsg{}

	jsonDec := json.NewDecoder(c.Request.Body)
	decer := jsonDec.Decode(in)
	if decer != nil {
		eres.Status = http.StatusBadRequest
		eres.Error = "invalid input data"
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "category_id", in.CategoryID)
	category, er, code := cathandler.CategorySer.GetCategoryByID(ctx)
	if (er != nil || code != state.DT_STATUS_OK) || category == nil {
		eres.Status = http.StatusNotFound
		eres.Error = "category with this ID not found"
		c.JSON(http.StatusNotFound, eres)
		return
	}
	if category.Fee == in.Amount {
		eres.Status = http.StatusNotModified
		eres.Error = "record was not modified"
		c.JSON(http.StatusNotModified, eres)
		return
	}
	category.Fee = in.Amount
	ctx = context.WithValue(ctx, "category", category)
	er, code = cathandler.CategorySer.UpdateCategory(ctx)
	if code == state.DT_STATUS_OK {
		c.JSON(http.StatusOK, category)
		return
	}
	println(er.Error())

	eres.Status = http.StatusInternalServerError
	eres.Error = "internal problem"
	c.JSON(http.StatusInternalServerError, eres)
}

func (cathandler *CategoryHandler) GetCategoryStudents(c *gin.Context) {
	ctx := c.Request.Context()
	eres := &model.ErMsg{}
	categoryID, er := strconv.Atoi(c.Query("category_id"))
	if er != nil {
		eres.Error = "bad request"
		eres.Status = http.StatusBadRequest
		c.JSON(eres.Status, eres)
		return
	}
	ctx = context.WithValue(ctx, "category_id", uint(categoryID))
	cat, status, er := cathandler.CategorySer.GetCategoryStudents(ctx)
	if cat != nil && status == state.DT_STATUS_OK && er != nil {
		c.JSON(http.StatusOK, cat)
		return
	}
	if cat == nil {
		eres.Error = " internal problem "
		eres.Status = http.StatusInternalServerError
		c.JSON(http.StatusNotFound, eres)
		return
	}
	c.JSON(http.StatusOK, cat)
}
