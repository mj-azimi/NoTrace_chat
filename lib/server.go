package lib

import (
	"NoTrace/config"
	"NoTrace/database/chats"

	"github.com/gin-gonic/gin"
	"github.com/inancgumus/screen"
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


		data := chats.Chat{
			ClientID: 100, 
			Text:     message, 
			ME:       false, 
		}
		chats.Create(data)
		screen.Clear()
		screen.MoveTopLeft()
		go ShowChat(1)

		ctx.JSON(200, gin.H{"status": "success"})	
	})
	server.Run(config.GetConfig("server_host"))
}
