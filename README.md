# 📦 Pack Distributor Application

[Pack Distributor Demo](https://cvele.github.io/reptask/) 

This repository contains an application that calculates optimal pack distributions for customer orders, minimizing waste and the number of packs used based on configurable pack sizes. It includes two main components: 

- **Backend (Go API)**: Manages pack sizes and calculates optimal pack distributions. 
- **Frontend (React)**: Provides a user-friendly interface to interact with the backend API. 

--- 

## Backend (Golang API) 

> Solution is deployed to **fly.io**, so expect a bit of latecy on initial request.

Core function `CalculateOptimalPacks` is [documented separatly](internal/pack/README.md).

### Features 

- **Pack Size Management**: Perform CRUD operations (Create, Read, Update, Delete) to manage available pack sizes. 
- **Optimal Distribution Calculation**: Implements dynamic programming to compute the minimum packs needed to efficiently fulfill customer orders. 

### Technology Stack 

- **Language**: Golang 
- **Database**: SQLite (recommended production options: PostgreSQL, DynamoDB, etc.) 
- **Containerization**: Docker (optimized multi-stage build) 

### Configuration Set environment variables for custom configuration: 

```bash 
export PORT=8080 export SQLITE_DB_PATH=./pack.db
``` 

### Docker Usage Build and run using Docker from the project's root directory: 

```bash 
make docker-build && make docker-run
``` 

The backend defaults to port `8080`. If occupied, specify another port with: 

```bash 
docker run -p <desired-port>:8080 <your-docker-image>
``` 

### Makefile Commands Convenient commands for managing your backend workflow:

```bash 
Available commands:
  make build         - Build the Go application
  make run           - Run the application
  make test          - Run unit tests
  make fmt           - Format code
  make vet           - Lint code
  make clean         - Clean build artifacts
  make docker-build  - Build Docker image
  make docker-run    - Run Docker container
  make docker-stop   - Stop Docker container
  make docker-test   - Build, run container, execute API tests, stop container
  make tidy          - Tidy up Go modules
  make deps          - Install dependencies
  make rebuild       - Clean, tidy, and build
  make prod-test     - Run curl tests against the production server
  make help          - Show this help message
``` 

---

## Frontend (React) 

The React frontend offers a ui experience to interact with the backend API, featuring: 
- **CRUD Interface**: Manage pack sizes easily. 
- **Order Calculation**: Calculate optimal pack distributions swiftly. 

### Running Frontend Locally To run the frontend locally: 

```bash 
cd frontend 
npm install && npm run dev
``` 

The frontend is hosted on GitHub Pages for easy accessibility. 

--- 

🎉 **Enjoy exploring the application, and I welcome your feedback!** 