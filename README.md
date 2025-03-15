# BYTEBITES
# Restaurant Food Ordering System

## Architecture Diagram
![Architecture Diagram](https://github.com/CallenCaracy/ByteBites/blob/main/documents/diagrams/SIA_Final_Project.drawio.png)

## Overview
The **Restaurant Food Ordering System** is a microservices-based architecture designed to handle restaurant operations efficiently. It includes menu management, order processing, payment handling, and kitchen operations, all communicating through gRPC internally and exposed via a GraphQL API layer.

## Tech Stack
- **Frontend:** React with Tailwind on Vite
- **Backend:** Go (GoLang)
- **Database:** PostgreSQL (Supabase) (PGX Driver)
- **API Layer:** GoQL (GraphQL)
- **Internal Communication:** gRPC
- **Database Migration:** Goose

## Microservices Architecture

### 1. **Menu Service**
   - Manages menu items and employee accounts.
   - Stores menu data, prices, and availability in the Menu DB.
   - Built with GoLang.

### 2. **Order Service**
   - Handles order placement and tracking.
   - Stores order data in the Order DB.
   - Interacts with Payment Service for payment processing.
   - Built with GoLang.

### 3. **Payment Service**
   - Processes payments securely.
   - Stores and retrieves payment records from the Payment DB.
   - Confirms payments for orders.
   - Built with GoLang.

### 4. **Kitchen Service**
   - Manages kitchen operations and order fulfillment.
   - Queries orders by first come, first serve.
   - Reads data from other databases for order tracking.
   - Built with GoLang.

### 5. **GraphQL DB Service**
   - Acts as the API layer for data aggregation.
   - Serves as the main gateway for frontend applications.
   - Uses GraphQL for efficient data querying.

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

