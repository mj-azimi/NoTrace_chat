package Chat

import (
	"database/sql"
	"log"
	"NoTrace/config"
)

type Chat struct {
	ID        int
	ME		  bool
	ClientID  int
	Text      string
	CreatedAt string
}

func Create(chat Chat) {
	db := config.Connect()
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO chats (client_id, text, me) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(chat.ClientID, chat.Text, chat.ME)
	if err != nil {
		log.Fatal(err)
	}

}


func FindById(id int) *Chat {
	db := config.Connect()
	defer db.Close()

	var chat Chat
	row := db.QueryRow("SELECT id, client_id, text, created_at FROM chats WHERE id = ?", id)
	err := row.Scan(&chat.ID, &chat.ClientID, &chat.Text, &chat.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil 
		} else {
			log.Fatal(err)
		}
	}
	return &chat
}


func All(client_id int) []Chat {
	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, me , client_id, text, created_at FROM chats WHERE client_id = ?", client_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var chats []Chat
	for rows.Next() {
		var chat Chat
		err := rows.Scan(&chat.ID,&chat.ME, &chat.ClientID, &chat.Text, &chat.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		chats = append(chats, chat)
	}

	return chats
}


func Delete(id int) {
	db := config.Connect()
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM chats WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

func Update(id int, newText string) {
	db := config.Connect()
	defer db.Close()

	statement, err := db.Prepare("UPDATE chats SET text = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(newText, id)
	if err != nil {
		log.Fatal(err)
	}
}
