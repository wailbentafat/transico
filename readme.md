Tansico 🏗️🔗
Tansico is a microservices-based platform that integrates Go for database and API handling, Express.js & Ethers.js for blockchain interactions, and Redis Pub/Sub for real-time messaging. The project uses SQLite for lightweight database storage.

📌 Overview
🔹 Components
Go Database Service – Manages SQLite-based database operations.
Express.js Blockchain Service – Handles blockchain interactions via Ethers.js.
Redis Pub/Sub – Enables communication between services.
SQLite – Local database for lightweight storage.
🛠 Tech Stack
Component	Technology
Backend API	Go (Gin/Fiber)
Blockchain Interaction	Node.js, Express.js, Ethers.js
Messaging	Redis Pub/Sub
Database	SQLite
Deployment	Docker, Docker Compose
📁 Project Structure
bash
Copy
Edit
/tansico
│── /databaseservice     # Go-based database service
│   ├── cmd              # Main entry point
│   ├── internal
│   │   ├── config       # Configuration files
│   │   ├── db           # Database connection & migrations
│   │   ├── handler      # API route handlers
│   │   ├── middleware   # Authentication, logging, etc.
│   │   ├── models       # Database models
│   │   ├── repository   # Database queries
│   ├── go.mod
│   ├── go.sum
│   ├── sqlite.db        # SQLite database file
│── /blockchain-service  # Express.js & Ethers.js service
│   ├── index.js
│   ├── contract.js
│   ├── redis_sub.js
│── /redis               # Redis configuration
│── docker-compose.yml   # Docker setup
│── README.md            # Documentation
🚀 Setup & Installation
1️⃣ Clone the Repository
bash
Copy
Edit
git clone https://github.com/yourusername/tansico.git
cd tansico
2️⃣ Install Go Modules
bash
Copy
Edit
cd databaseservice
go mod tidy
3️⃣ Run SQLite Migrations
Ensure SQLite is installed. If not, install it:

bash
Copy
Edit
sudo apt install sqlite3  # Ubuntu/Debian
brew install sqlite3      # macOS
Initialize the database:

bash
Copy
Edit
sqlite3 sqlite.db < internal/db/schema.sql
4️⃣ Start Redis (Standalone Mode)
If Redis isn’t installed, run it using Docker:

bash
Copy
Edit
docker run --name redis -p 6379:6379 -d redis
5️⃣ Run the Go Database Service
bash
Copy
Edit
cd databaseservice
go run cmd/main.go
6️⃣ Run the Blockchain Service (Express.js & Ethers.js)
bash
Copy
Edit
cd blockchain-service
npm install
node index.js
📡 How It Works (SQLite & Redis Pub/Sub Flow)
1️⃣ The Go service provides REST API endpoints to store and retrieve data from SQLite.
2️⃣ When blockchain-related data needs processing, the service publishes messages to Redis.
3️⃣ The Express.js blockchain service listens to Redis, interacts with the blockchain, and updates the database accordingly.

🔄 API Endpoints (Go REST API)
Method	Endpoint	Description
POST	/data	Stores data in SQLite
GET	/data/:id	Retrieves data from SQLite
GET	/status	Checks API health
Example Request:

bash
Copy
Edit
curl -X POST http://localhost:3000/data -H "Content-Type: application/json" -d '{"key": "value"}'
📌 SQLite Database Schema
Example schema.sql for initializing SQLite:

sql
Copy
Edit
CREATE TABLE IF NOT EXISTS records (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key TEXT NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
📌 Future Improvements
✅ Add WebSocket notifications for real-time updates.
✅ Implement JWT authentication for API security.
✅ Support multiple blockchain networks.

🤝 Contributing
Pull requests are welcome! Fork the repo, create a feature branch, and submit a PR.

