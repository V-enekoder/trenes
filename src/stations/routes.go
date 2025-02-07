package station

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	stations := router.Group("/stations")
	{
		stations.POST("/", CreateStationController)
		stations.GET("/:id", GetStationByIDController)
		stations.GET("/", GetAllStationsController)
		stations.PUT("/:id", UpdateStationController)
		stations.DELETE("/:id", DeleteStationController)
		stations.GET("/ruta-optima/:start_id/:end_id", FindOptimalRoadController)
	}
}
