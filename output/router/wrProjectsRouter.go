package router

import (
	. "pm-system/controllers"
	"github.com/gin-gonic/gin"
)

func wrProjectsRouter(router *gin.Engine) {
	wrProjects := WrProjectsController{}
	wrProjectsR := router.Group("projects")
	{
		wrProjectsR.GET("/", wrProjects.Find)
		wrProjectsR.GET("/one", wrProjects.FindOne)
		wrProjectsR.GET("/page", wrProjects.FindPaging)
		wrProjectsR.POST("/", wrProjects.Create)
		wrProjectsR.GET("/find-all", wrProjects.FindAll)
		wrProjectsR.GET("/find-by-id/:id", wrProjects.FindById)
		wrProjectsR.PUT("/update-by-id", wrProjects.UpdateById)
	}
}