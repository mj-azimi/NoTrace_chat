package main

import (
	"fmt"
	"gitlab.com/pragmaticreviews/golang-gin-poc/lib"
)


 func main()  {
	go lib.Server()
	
	data := map[string]interface{}{
		"message": "hello how are you",
	}
	res , _ :=  lib.SendRequest("http://localhost:8999/get_text",data)

	fmt.Println(res)
 }