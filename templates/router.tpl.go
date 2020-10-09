package templates

var RouterTpl = `package router

import (
	. "gpi/controllers"
	"github.com/gin-gonic/gin"
)

func {{StructFcName}}Router(router *gin.Engine) {
	{{StructFcName}} := {{StructName}}Controller{}
	{{StructFcName}}R := router.Group("{{StructRoute}}")
	{
		{{StructFcName}}R.GET("/", {{StructFcName}}.Find)
		{{StructFcName}}R.GET("/page", {{StructFcName}}.FindPaging)
		{{StructFcName}}R.POST("/", {{StructFcName}}.Create)
		{{StructFcName}}R.GET("/find-by-id/:id", {{StructFcName}}.FindById)
		{{StructFcName}}R.PUT("/update-by-id", {{StructFcName}}.UpdateById)
	}
}`
