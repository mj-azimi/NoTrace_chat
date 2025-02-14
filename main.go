package main

import (
	"NoTrace/bootstrap"
	"NoTrace/lib"

)

func main() {
	bootstrap.Boot()
	go lib.Server()
	lib.ShowChat()
	
}
