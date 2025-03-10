# BYTEBITES
# Restaurant Food Ordering System

## Overview
The **Restaurant Food Ordering System** is a microservices-based architecture designed to handle restaurant operations efficiently. It includes menu management, order processing, payment handling, and kitchen operations, all communicating through gRPC internally and exposed via a GraphQL API layer.

## Tech Stack
- **Frontend:** React with Tailwind on Vite
- **Backend:** Go (GoLang)
- **Database:** PostgreSQL (Supabase)
- **API Layer:** GoQL (GraphQL)
- **Internal Communication:** gRPC
- **Database Migration:** Goose

## Microservices Architecture

### 1. **Menu Service**
   - Manages menu items
   - Built with GoLang and GraphQL

### 2. **Order Service**
   - Handles order placement
   - Built with GoLang and GraphQL

### 3. **Payment Service**
   - Processes payments securely
   - Built with GoLang and GraphQL

### 4. **Kitchen Service**
   - Manages kitchen operations and order fulfillment
   - Built with GoLang and GraphQL

### 5. **GraphQL DB Service**
   - Provides a GraphQL API for database operations
   - Uses PostgreSQL (Supabase) as the database

## Installation & Setup

### Prerequisites
- [Go](https://go.dev/dl/)
- [Node.js](https://nodejs.org/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Supabase](https://supabase.com/)
- [Goose](https://github.com/pressly/goose)

### Clone the Repository
```sh
git clone https://github.com/CallenCaracy/ByteBites.git
cd ByteBites
```

### Database Setup
1. Set up a **PostgreSQL** database using **Supabase**.
2. Configure database credentials in the environment file.
3. Run database migrations using Goose:
   ```sh
   goose up
   ```

### Backend Setup
1. Navigate to the backend directory:
   ```sh
   cd backend
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Start microservices:
   ```sh
   go run main.go
   ```

### Frontend Setup
1. Navigate to the frontend directory:
   ```sh
   cd frontend
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
3. Start the React app:
   ```sh
   npm start
   ```

## API Documentation
- **GraphQL API Endpoint:** `http://localhost:4000/graphql`
- **gRPC Communication:** Internal microservices use gRPC for inter-service communication.

## Contributing
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit changes (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a pull request.

## License
This project is licensed under the MIT License.

---
Feel free to customize this README as per your project needs!

