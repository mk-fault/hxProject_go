package routers

import (
	"hx/controller"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("api")
	{
		statsGroup := apiGroup.Group("stats")
		{
			statsGroup.GET("/qg", controller.GetStatsQGList)
			statsGroup.GET("/gs", controller.GetStatsGSList)
		}
		whoGroup := apiGroup.Group("who")
		{
			whoGroup.GET("/", controller.GetWHOList)
			whoGroup.GET("/index", controller.GetWHOIndex)
		}
		gbdGroup := apiGroup.Group("gbd")
		{
			gbdGroup.GET("/", controller.GetGBDList)
		}
	}

	return r
}