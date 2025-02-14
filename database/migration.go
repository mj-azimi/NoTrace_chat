package database

import "NoTrace/config"

func Handle(){
	clients()
	chats()
}

func clients()(bool) {
	db := config.Connect();
	defer db.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		ip TEXT
	);`
	db.Exec(createTableSQL)
	return true;
}

func chats()(bool) {
	db := config.Connect();
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
	return true;
}