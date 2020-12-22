package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pjchender/go-snippets/template/model"
)

// ProductDatabase 包含在 ProductAPI 中所有會使用到的 Database 方法
type ProductDatabase interface {
	CreateProduct(product *model.Product) error
	GetProducts() ([]*model.Product, error)
	GetProductsWithConditions(conditions ...interface{}) ([]*model.Product, error)
	GetProductByID(productID uuid.UUID) (*model.Product, error)
	UpdateProductWithZero(product *model.Product) error
	UpsertProductByProviderWithZero(product *model.Product) (*model.Product, error)
	DeleteProductByID(productID uuid.UUID) error
}

// ProductAPI 可以使用 ProductDatabase 的方法
type ProductAPI struct {
	DB ProductDatabase
}

// CreateProduct 會建立 Product
func (p *ProductAPI) CreateProduct(ctx *gin.Context) {
	var err error
	var product model.Product
	err = ctx.Bind(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = p.DB.CreateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// GetProducts 會回傳所有的 Products
func (p *ProductAPI) GetProducts(ctx *gin.Context) {
	products, err := p.DB.GetProducts()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// GetProductsWithConditions 可以透過 queryString 篩選使用者想要得資料
func (p *ProductAPI) GetProductsWithConditions(ctx *gin.Context) {
	var qs model.ProductQuery
	err := ctx.BindQuery(&qs)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	products, err := p.DB.GetProductsWithConditions(&model.Product{
		Name:      qs.Name,
		IsPublish: strToBool(qs.IsPublish),
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

	ctx.JSON(http.StatusOK, product)
}

// UpdateProductByID 會根據使用者輸入的內容更新 product
func (p *ProductAPI) UpdateProductByID(ctx *gin.Context) {
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

	product := model.Product{
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

func strToBool(str string) bool {
	if str == "true" {
		return true
	}

	return false
}

func toExternalProducts(products []*model.Product) []*model.ProductExternal {
	productsExternal := make([]*model.ProductExternal, len(products))

	for i, product := range products {
		productsExternal[i] = product.ToExternal()
	}

	return productsExternal
}
