package core

import (
	"github.com/shop-market/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Router(connection *gorm.DB) {
	router := gin.Default()

	// A good base middleware stack
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Set a timeout value on the request context (ctx), that will signal

	// RESTy routes for "auth" resource
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			handler := api_v1.Authority(connection)

			auth.POST("/register", handler.Register) // POST /auth/register
		}
	}

	router.Run(":3001")
}
