package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRoute(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	return router
}
