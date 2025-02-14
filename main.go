package main

import (
	// "fmt"
	// "gitlab.com/pragmaticreviews/golang-gin-poc/lib"
	"NoTrace/bootstrap"
	Chat "NoTrace/model"
	// "NoTrace/lib"
)


 func main()  {
	bootstrap.Boot()

	newChat := Chat.Chat{
		ClientID: 1, // شناسه مشتری
		Text:     "سلام، چطور می‌توانم به شما کمک کنم؟", // متن چت
		ME:       true, // مقدار برای فیلد ME (true یا false)
	}

	// فراخوانی تابع Create برای افزودن چت جدید به دیتابیس
	Chat.Create(newChat)

	// go lib.Server()
	
	// data := map[string]interface{}{
	// 	"message": "hello how are you",
	// }
	// res , _ :=  lib.SendRequest("http://localhost:8999/get_text",data)

	// fmt.Println(res)
 }