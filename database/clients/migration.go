package clients

import "NoTrace_chat/config"

func Migration() {
	db := config.Connect()
	defer db.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		ip TEXT
	);`
	db.Exec(createTableSQL)
}
