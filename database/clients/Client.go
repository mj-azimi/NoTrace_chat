package clients

import (
	"database/sql"
	"log"
	"NoTrace/config"
)

type Client struct {
	ID   int
	Name string
	IP   string
}

func Create(client Client) {
	db := config.Connect()
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO clients (name, ip) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(client.Name, client.IP)
	if err != nil {
		log.Fatal(err)
	}
}

func FindById(id int) *Client {
	db := config.Connect()
	defer db.Close()

	var client Client
	row := db.QueryRow("SELECT id, name, ip FROM clients WHERE id = ?", id)
	err := row.Scan(&client.ID, &client.Name, &client.IP)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		} else {
			log.Fatal(err)
		}
	}
	return &client
}

func All() []Client {
	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, ip FROM clients")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var client Client
		err := rows.Scan(&client.ID, &client.Name, &client.IP)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, client)
	}

	return clients
}

func Delete(id int) {
	db := config.Connect()
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM clients WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

func Update(id int, newName string, newIP string) {
	db := config.Connect()
	defer db.Close()

	statement, err := db.Prepare("UPDATE clients SET name = ?, ip = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(newName, newIP, id)
	if err != nil {
		log.Fatal(err)
	}
}
