package bootstrap

import "NoTrace/database"

func Boot()  {
	database.Handle()
}