package lib

import (
	"NoTrace/database/chats"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func SendRequest(url string, jsonData map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	message := jsonData["message"].(string)

	data := chats.Chat{
		ClientID: 100, 
		Text:     message, 
		ME:       true, 
	}
	chats.Create(data)

	return string(body), nil
}
