![Image](https://s6.uupload.ir/files/screencast_from_25-02-15_14_56_04_mvwy.gif)

# NoTrace - Chat System Based on Go and SQLite

NoTrace is a lightweight chat system developed using **Go**, **Gin**, and **SQLite**. This project includes a **command-line interface** for sending and receiving messages, as well as a **RESTful API server** for managing chats.


## new Features
- Ability to save contacts
- Chat selection between brains
- Fix bugs

---

##  Privacy & Security

NoTrace_chat does **not store chats on any external server**, making it **highly secure**. All messages are stored **locally** using SQLite, ensuring that no intermediary servers are needed for communication.

---

##  Features

- Store and display messages using **SQLite**.
- Command-line interface for message exchange.
- **RESTful API** server using **Gin** framework.
- Messages stored with timestamps.
- Ability to delete and update messages.
- **No third-party server involvement** for privacy and security.

---

##  Installation and Usage

### Prerequisites
- **Go** (version 1.19 or later).

### Steps to Run the Project
1. Clone the repository:
   ```sh
   git clone https://github.com/mj-azimi/NoTrace_chat.git
   cd NoTrace_chat
   ```
2. Install dependencies:

    ```
    go mod tidy
    ```
3. Build the project:
    ```
        go build
    ```
4. And run App
    ```
        ./NoTrace_chat
    ```

##  Project Structure
    
        NoTrace_chat/
        ├── bootstrap/       # Database initialization
        ├── config/          # Configuration management and database connection
        ├── database/        # SQLite database handling
        ├── lib/             # Server and CLI management
        ├── model/           # Data model for chats
        ├── main.go          # Entry point of the application
        ├── go.mod           # Go dependency management
        ├── go.sum           # Dependency lock file
        ├── config.json      # Configuration file
        └── README.md        # This file
    
###  License
This project is licensed under the MIT License.

##  Contact

For suggestions or issues, please open an Issue on the GitHub repository!
