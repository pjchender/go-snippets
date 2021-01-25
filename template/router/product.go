package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/go-snippets/template/database"
	v1 "github.com/pjchender/go-snippets/template/router/api/v1"
)

// RegisterProductV1 用來建立和 Product 有關的路由
func RegisterProductV1(db *database.GormDatabase, routerGroup *gin.RouterGroup) {
	productHandler := v1.NewProductHandler(db)
	productRouter := routerGroup.Group("/products")
	{
		productRouter.GET("/", productHandler.GetProducts)
		productRouter.POST("/", productHandler.CreateProduct)
		productRouter.GET("/:id", productHandler.GetProductByID)
		productRouter.PATCH("/:id", productHandler.UpdateProductByID)
		productRouter.DELETE("/:id", productHandler.DeleteProductByID)
	}
}
