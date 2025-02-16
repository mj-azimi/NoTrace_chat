package database

import (
	"NoTrace/database/chats"
	"NoTrace/database/clients"
)

func Handle(){
	clients.Migration()
	chats.Migration()
}

