package lib

import (
	"github.com/gin-gonic/gin"
)

func Server() {
	server := gin.Default()

	server.POST("/get_text", func(ctx *gin.Context) {
		var jsonData map[string]interface{}
		// دریافت داده‌های JSON از درخواست
		if err := ctx.ShouldBindJSON(&jsonData); err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		
		// پاسخ با پیام OK
		ctx.JSON(200, gin.H{
			"mess": "ok",
		})
	})

	server.Run(":8999")
}
