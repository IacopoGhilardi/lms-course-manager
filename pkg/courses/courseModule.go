package courses

import (
	coursesControllers "github.com/IacopoGhilardi/lms-course-manager/pkg/courses/controllers"
	"github.com/gin-gonic/gin"
)

func InitModule(router *gin.RouterGroup) {
	courseRouter := router.Group("/courses")

	coursesControllers.InitRoutes(*courseRouter)
}
