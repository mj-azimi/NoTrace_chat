package database

import (
	"NoTrace_chat/database/chats"
	"NoTrace_chat/database/clients"
)

func Handle() {
	clients.Migration()
	chats.Migration()
}
