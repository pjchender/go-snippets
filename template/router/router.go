package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/go-snippets/template/database"
)

// NewRouter 用來建立所有的 router
func NewRouter(db *database.GormDatabase) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		RegisterProductV1(db, v1)
	}

	return router
}
