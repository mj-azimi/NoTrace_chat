package lib

import (
	Chat "NoTrace/model"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShowChat()  {
	for {
		chats := Chat.All(100)
		for _, chat := range chats {
			if chat.ME {
				fmt.Printf("* me ⚪ *\n%s\n****\n%s\n", chat.Text, chat.CreatedAt)
			} else {
				fmt.Printf("* 👤 *\n%s\n%s\n", chat.Text, chat.CreatedAt)
			}
			fmt.Println("---------------------")
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a message: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting...")
			break
		}

		data := map[string]interface{}{
			"message": text,
		}
		SendRequest("http://localhost:7999/get_text", data)
	}
}