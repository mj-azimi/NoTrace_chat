# NoTrace - Chat System Based on Go and SQLite

NoTrace is a lightweight chat system developed using **Go**, **Gin**, and **SQLite**. This project includes a **command-line interface** for sending and receiving messages, as well as a **RESTful API server** for managing chats.

---

## 🔒 Privacy & Security

NoTrace does **not store chats on any external server**, making it **highly secure**. All messages are stored **locally** using SQLite, ensuring that no intermediary servers are needed for communication.

---

## ✨ Features

- Store and display messages using **SQLite**.
- Command-line interface for message exchange.
- **RESTful API** server using **Gin** framework.
- Messages stored with timestamps.
- Ability to delete and update messages.
- **No third-party server involvement** for privacy and security.

---

## 🚀 Installation and Usage

### Prerequisites
- **Go** (version 1.19 or later).

### Steps to Run the Project
1. Clone the repository:
   ```sh
   git clone https://github.com/mj-azimi/NoTrace.git
   cd NoTrace
   ```
2. Install dependencies:

    ```
    go mod tidy
    ```
3. Build the project:
    ```
        go build
    ```
4. Adn run App
    ```
        ./NoTrace
    ```

## 🔧 Project Structure
    
        NoTrace/
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
    
### 📜 License
This project is licensed under the MIT License.

## 📞 Contact

For suggestions or issues, please open an Issue on the GitHub repository!