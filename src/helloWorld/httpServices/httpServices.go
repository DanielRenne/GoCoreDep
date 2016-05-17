package httpServices

import (
	"fmt"
	"github.com/DanielRenne/GoCore/core/ginServer"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("helloWorld httpServices initialized.")

	ginServer.Router.GET("WebAPI", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
