package routes

import (
	"go-chatgpt/controllers"

	"github.com/gin-gonic/gin"
)

func ChatGPTRoutes(r *gin.RouterGroup) {
	r.POST("/chat-gpt", controllers.GenerateChat) // Generate chat
}
