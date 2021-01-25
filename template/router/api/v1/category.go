package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pjchender/go-snippets/template/database"
	"github.com/pjchender/go-snippets/template/model"
	"github.com/pjchender/go-snippets/template/service"
	log "github.com/sirupsen/logrus"
)

// CategoryAPI  會將呼叫 service 需要取得的參數從 HTTP Request 中取出後
// 呼叫在 service 中 *database.GormDatabase 的方法，
// 並處理 Bad reqest 等類型的錯誤
type CategoryAPI struct {
	DB *database.GormDatabase
}

// NewCategoryHandler 會回傳 *CategoryAPI，以此使用 *CategoryAPI 建立的方法
func NewCategoryHandler(db *database.GormDatabase) *CategoryAPI {
	return &CategoryAPI{
		DB: db,
	}
}

func (c *CategoryAPI) Get(ctx *gin.Context) {
	var err error
	categoryIDStr := ctx.Param("id")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	svc := service.New(ctx, c.DB)
	category, err := svc.GetCategoryByID(service.GetCategoryRequest{ID: categoryID})

	if err != nil {
		log.Errorf("svc.GetCategoryByID failed: %v", err)
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, category.ToExternal())
}

func (s *CategoryAPI) List(ctx *gin.Context) {
	param := service.ListCategoryRequest{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		log.Errorf("ctx.ShouldBind failed: %v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	svc := service.New(ctx.Request.Context(), s.DB)
	categories, err := svc.GetCategoryList(param)
	if err != nil {
		log.Errorf("svc.GetCategory failed: %v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, model.ToExternalCategory(categories))
}

func (c *CategoryAPI) Create(ctx *gin.Context) {
	param := service.CreateCategoryRequest{}
	err := ctx.ShouldBind(&param)
	if err != nil {
		log.Errorf("ctx.ShouldBind failed: %v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	svc := service.New(ctx.Request.Context(), c.DB)
	category, err := svc.CreateCategory(param)
	if err != nil {
		log.Errorf("svc.CreateCategory failed: %v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, category.ToExternal())
}

func (c *CategoryAPI) BatchCreate(ctx *gin.Context) {
	params := []*service.CreateCategoryRequest{}
	err := ctx.ShouldBind(&params)
	if err != nil {
		log.Errorf("ctx.ShouldBind failed: %v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	svc := service.New(ctx.Request.Context(), c.DB)
	categories, err := svc.CreateCategories(params)
	if err != nil {
		log.Errorf("svc.CreateCategories failed: %v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, model.ToExternalCategory(categories))
}

func (c *CategoryAPI) Update(ctx *gin.Context) {
	var err error
	categoryIDStr := ctx.Param("id")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		log.Errorf("uuid.Parse failed: %v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	param := service.UpdateCategoryRequest{ID: categoryID}
	err = ctx.ShouldBind(&param)
	if err != nil {
		log.Errorf("ctx.ShouldBind failed: %v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	svc := service.New(ctx.Request.Context(), c.DB)

	updatedCategory, err := svc.UpdateCategory(param)
	if err != nil {
		log.Errorf("svc.UpdateCategory failed: %v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedCategory.ToExternal())
}

func (c *CategoryAPI) Delete(ctx *gin.Context) {
	var err error
	categoryIDStr := ctx.Param("id")
	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		log.Errorf("uuid.Parse failed: %v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	svc := service.New(ctx.Request.Context(), c.DB)

	err = svc.DeleteCategory(service.DeleteCategoryRequest{
		ID: categoryID,
	})
	if err != nil {
		log.Errorf("svc.DeleteCategory failed: %v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
