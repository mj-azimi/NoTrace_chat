package chats

import "NoTrace_chat/config"

func Migration() {
	db := config.Connect()
	defer db.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS chats (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		client_id INTEGER NOT NULL,
		me BOOLEAN,
		text TEXT ,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	db.Exec(createTableSQL)
}
