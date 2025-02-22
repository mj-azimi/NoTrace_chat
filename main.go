package main

import (
	"NoTrace/bootstrap"
	"NoTrace/database/clients"
	"NoTrace/lib"
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/manifoldco/promptui"
)

func main() {
	for {
		screen.MoveTopLeft()

		bootstrap.Boot()

		prompt := promptui.Select{
			Label: "New chat or contacts?",
			Items: []string{"Contacts 👤", "Add Contacts 👤"},
		}

		num, _, err := prompt.Run()
		if err != nil {
			fmt.Println("❌ Error receiving input :", err)
			return
		}

		if num == 0 {
			clientsData := clients.All()
			var cl []string
			clientIDMap := make(map[int]int) 

			for index, client := range clientsData {
				cl = append(cl, fmt.Sprintf("%s *** %s", client.Name, client.IP))
				clientIDMap[index] = client.ID
			}

			prompt := promptui.Select{
				Label: "New chat or contacts?",
				Items: cl,
			}

			num, _, err := prompt.Run()
			if err != nil {
				fmt.Println("❌ Error receiving input:", err)
				return
			}

			selectedID := clientIDMap[num]

			go lib.Server()
			lib.ShowChat(selectedID)
		} else {

			var ip, client string

			fmt.Println("enter new ip and port: \n (sample: http://127.0.0.1:1234)")
			fmt.Scanln(&ip)

			fmt.Println("Enter the user name:")
			fmt.Scanln(&client)

			data := clients.Client{
				Name: client,
				IP:   ip,
			}
			clients.Create(data)

			fmt.Println("Saved successfully. ✅")

			time.Sleep(2 * time.Second)
			screen.Clear()
			screen.MoveTopLeft()

		}
		continue

	}

}
