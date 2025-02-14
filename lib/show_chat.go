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
		// نمایش پیام‌های قبلی
		chats := Chat.All(100)
		for _, chat := range chats {
			if chat.ME {
				fmt.Printf("* me ⚪ *\n%s\n****\n%s\n", chat.Text, chat.CreatedAt)
			} else {
				fmt.Printf("* 👤 *\n%s\n%s\n", chat.Text, chat.CreatedAt)
			}
			fmt.Println("---------------------")
		}

		// دریافت ورودی از کاربر
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a message: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		// اگر کاربر "exit" تایپ کرد، برنامه متوقف شود
		if text == "exit" {
			fmt.Println("Exiting...")
			break
		}

		// ارسال درخواست
		data := map[string]interface{}{
			"message": text,
		}
		SendRequest("http://localhost:7999/get_text", data)
	}
}