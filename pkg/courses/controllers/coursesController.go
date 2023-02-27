package coursesControllers

import "github.com/gin-gonic/gin"

func InitRoutes(router gin.RouterGroup) {

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ciao",
		})
	})
}
