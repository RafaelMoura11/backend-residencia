package routes

import (
	"github.com/gin-gonic/gin"
	"backend-residencia/controllers"
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

	message := r.Group("/messages")
	{
		message.POST("/", controllers.CreateMessage)
		message.GET("/", controllers.GetMessages)
		message.GET("/:id", controllers.GetMessageByID)
		message.PUT("/:id", controllers.UpdateMessage)
		message.DELETE("/:id", controllers.DeleteMessage)
	}

	usage := r.Group("/usage-tokens")
	{
		usage.POST("/", controllers.CreateUsageToken)
		usage.GET("/", controllers.GetUsageTokens)
		usage.GET("/:id", controllers.GetUsageTokenByID)
		usage.PUT("/:id", controllers.UpdateUsageToken)
		usage.DELETE("/:id", controllers.DeleteUsageToken)
	}

	tool := r.Group("/tools")
	{
		tool.POST("/", controllers.CreateTool)
		tool.GET("/", controllers.GetTools)
		tool.GET("/:id", controllers.GetToolByID)
		tool.PUT("/:id", controllers.UpdateTool)
		tool.DELETE("/:id", controllers.DeleteTool)
	}

	return r
}
