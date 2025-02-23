package main

import (
	"NoTrace_chat/bootstrap"
	"NoTrace_chat/config"
	"NoTrace_chat/database/clients"
	"NoTrace_chat/lib"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/inancgumus/screen"
	"github.com/manifoldco/promptui"
)

func main() {

	for {
		screen.MoveTopLeft()


		out, err := exec.Command("hostname", "-I").Output()
		if err != nil {
			fmt.Println("err :", err)
			return
		}
		println("\n \n \a");

		ips := strings.Fields(string(out))
		fmt.Println("🌐 IPs:", ips)

		fmt.Println("📡 PORT :", config.GetConfig("server_host") )
		
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
