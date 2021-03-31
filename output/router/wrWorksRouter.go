package router

import (
	. "pm-system/controllers"
	"github.com/gin-gonic/gin"
)

func wrWorksRouter(router *gin.Engine) {
	wrWorks := WrWorksController{}
	wrWorksR := router.Group("works")
	{
		wrWorksR.GET("/", wrWorks.Find)
		wrWorksR.GET("/one", wrWorks.FindOne)
		wrWorksR.GET("/page", wrWorks.FindPaging)
		wrWorksR.POST("/", wrWorks.Create)
		wrWorksR.GET("/find-all", wrWorks.FindAll)
		wrWorksR.GET("/find-by-id/:id", wrWorks.FindById)
		wrWorksR.PUT("/update-by-id", wrWorks.UpdateById)
	}
}