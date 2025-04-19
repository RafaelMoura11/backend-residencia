package routes

import (
	"github.com/gin-gonic/gin"
	"your_project/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	agent := r.Group("/agents")
	{
		agent.POST("/", controllers.CreateAgent)
		agent.GET("/", controllers.GetAgents)
		agent.GET("/:id", controllers.GetAgentByID)
		agent.PUT("/:id", controllers.UpdateAgent)
		agent.DELETE("/:id", controllers.DeleteAgent)
	}

	return r
}
