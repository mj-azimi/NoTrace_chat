package bootstrap

import "NoTrace_chat/database"

func Boot() {
	database.Handle()
}
