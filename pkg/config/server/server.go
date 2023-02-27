package server

import (
	"github.com/IacopoGhilardi/lms-course-manager/pkg/courses"
	"github.com/gin-gonic/gin"
)

func Init() {
	app := gin.Default()

	v1 := app.Group("/v1")

	courses.InitModule(v1)

	app.Run(":5000")
}
