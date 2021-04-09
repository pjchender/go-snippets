package v1

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pjchender/go-snippets/template/database"
	"github.com/pjchender/go-snippets/template/model"
	"github.com/pjchender/go-snippets/template/pkg/convert"
	"github.com/pjchender/go-snippets/template/service"
)

// ProductAPI 可以使用 ProductDatabase 的方法
type ProductAPI struct {
	DB *database.GormDatabase
}

// NewProductHandler 是用來建立 ProductAPI 這個 struct
func NewProductHandler(db *database.GormDatabase) *ProductAPI {
	return &ProductAPI{
		DB: db,
	}
}

// ProductQuery 提供可以在 url 後帶入使用的 queryString
type ProductQuery struct {
	ProductID string `form:"productId"`
	Name      string `form:"name"`
	IsPublish string `form:"isPublish"`

	// categoryID 雖然是 uuid 但透過 queryString 傳的時候只能是字串
	CategoryID string `form:"categoryId"`

	// query 時間的話可以用 int64 後續轉成 Unix
	CreatedAt int64 `form:"createdAt"`
	BeginDate int64 `form:"beginDate"`
	EndDate   int64 `form:"endDate"`
}

// CreateProduct 會建立 Product
func (p *ProductAPI) CreateProduct(ctx *gin.Context) {
	var err error

	var param service.CreateProductRequest
	err = ctx.Bind(&param)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isPublish, err := strconv.ParseBool(param.IsPublish)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	categoryID, err := uuid.Parse(param.CategoryID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product := model.Product{
		Name:       param.Name,
		Price:      param.Price,
		IsPublish:  isPublish,
		CategoryID: categoryID,
	}

	err = p.DB.CreateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product.ToExternal())
}

// GetProducts 會回傳所有的 Products
func (p *ProductAPI) GetProducts(ctx *gin.Context) {
	products, err := p.DB.GetProducts()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, toExternalProducts(products))
}

// GetProductsWithConditions 可以透過 queryString 篩選使用者想要得資料
func (p *ProductAPI) GetProductsWithConditions(ctx *gin.Context) {
	var param ProductQuery
	err := ctx.BindQuery(&param)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var isPublish bool
	if param.IsPublish != "" {
		isPublish, err = strconv.ParseBool(param.IsPublish)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	var productID uuid.UUID
	if param.ProductID != "" {
		productID, err = uuid.Parse(param.ProductID)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	var categoryID uuid.UUID
	if param.CategoryID != "" {
		categoryID, err = uuid.Parse(param.CategoryID)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	// query 特定某一天的時間
	var createdAt time.Time
	if param.CreatedAt != 0 {
		createdAt = time.Unix(param.CreatedAt, 0)
	}
	beginDate, endDate := convert.ParseTimeRange(param.BeginDate, param.EndDate)

	svc := service.New(ctx, p.DB)
	products, err := svc.ListProducts(
		service.ListProductRequest{
			ProductID:  productID,
			Name:       param.Name,
			IsPublish:  isPublish,
			CategoryID: categoryID,
			CreatedAt:  createdAt,
			BeginDate:  beginDate,
			EndDate:    endDate,
		})
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, toExternalProducts(products))
}

// GetProductByID 會根據路由中的 param :id 來找出對應的 product
func (p *ProductAPI) GetProductByID(ctx *gin.Context) {
	productIDStr := ctx.Param("id")

	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	product, err := p.DB.GetProductByID(productID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product.ToExternal())
}

// GetProductsInCategoryIDs 會根據 categoryIDs 回傳對應的 products
func (p *ProductAPI) GetProductsInCategoryIDs(ctx *gin.Context) {
	var qs struct {
		categoryIDs string `form:"categoryIds"`
	}

	err := ctx.BindQuery(&qs)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	categoryIDsStr := strings.Split(qs.categoryIDs, ",")
	categoryIDs := make([]uuid.UUID, len(categoryIDsStr))
	for i, categoryIDStr := range categoryIDsStr {
		productID, err := uuid.Parse(categoryIDStr)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		categoryIDs[i] = productID
	}

	products, err := p.DB.GetProductsInCategoryIDs(categoryIDs)

	ctx.JSON(http.StatusOK, toExternalProducts(products))
}

// UpdateProductByID 會根據使用者輸入的內容更新 product
func (p *ProductAPI) UpdateProductByID(ctx *gin.Context) {
	var err error
	productIDStr := ctx.Param("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 先確認該筆 record 存在，找不到該 record 則回傳 404
	// 如果沒先確認該筆 record 存在的話，即使沒有該 record 也會回傳 200
	_, err = p.DB.GetProductByID(productID)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	var productExternal model.ProductExternal
	err = ctx.Bind(&productExternal)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product := service.UpdateProductRequest{
		ID:        productID,
		Name:      productExternal.Name,
		Price:     productExternal.Price,
		IsPublish: productExternal.IsPublish,
	}
	err = p.DB.UpdateProductWithZero(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// DeleteProductByID 會根據路由中 param :id 來刪除特定的 product
func (p *ProductAPI) DeleteProductByID(ctx *gin.Context) {
	productIDStr := ctx.Param("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 先確認該筆 record 存在，找不到該 record 則回傳 404
	// 如果沒先確認該筆 record 存在的話，即使沒有該 record 也會回傳 200
	_, err = p.DB.GetProductByID(productID)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	err = p.DB.DeleteProductByID(productID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func toExternalProducts(products []*model.Product) []*model.ProductExternal {
	productsExternal := make([]*model.ProductExternal, len(products))

	for i, product := range products {
		productsExternal[i] = product.ToExternal()
	}

	return productsExternal
}
