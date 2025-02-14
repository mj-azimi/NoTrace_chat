package lib

import (
	"NoTrace/config"
	// Chat "NoTrace/model"

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
		
		// data := Chat.Chat{

		// };
		// Chat.Create(data)
	})
	server.Run(config.GetConfig("server_host"))
}
