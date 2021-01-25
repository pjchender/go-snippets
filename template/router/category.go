package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/go-snippets/template/database"
	v1 "github.com/pjchender/go-snippets/template/router/api/v1"
)

// RegisterCategory create route for Category
func RegisterCategory(db *database.GormDatabase, routerGroup *gin.RouterGroup, middleware ...gin.HandlerFunc) {
	NewCategoryHandler := v1.NewCategoryHandler(db)
	categoryRouter := routerGroup.Group("/categories", middleware...)
	{
		categoryRouter.GET("/", NewCategoryHandler.List)
		categoryRouter.GET("/:id", NewCategoryHandler.Get)
		categoryRouter.POST("/", NewCategoryHandler.Create)
		categoryRouter.PATCH("/:id", NewCategoryHandler.Update)
		categoryRouter.DELETE("/:id", NewCategoryHandler.Delete)
	}
}
