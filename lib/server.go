package lib

import (
	"NoTrace/config"
	Chat "NoTrace/model"

	"github.com/gin-gonic/gin"
)

func Server() {
	server := gin.Default()

	server.POST("/get_text", func(ctx *gin.Context) {
		var jsonData map[string]interface{}
		if err := ctx.ShouldBindJSON(&jsonData); err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid JSON",
			})
			return
		}
		
		message := jsonData["message"].(string)


		data := Chat.Chat{
			ClientID: 100, 
			Text:     message, 
			ME:       false, 
		}
		Chat.Create(data)
		go ShowChat()

		ctx.JSON(200, gin.H{"status": "success"})	})
	server.Run(config.GetConfig("server_host"))
}
