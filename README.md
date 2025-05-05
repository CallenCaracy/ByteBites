# BYTEBITES - Restaurant Food Ordering System

## Overview
**ByteBites** is a **microservices-based** Restaurant Food Ordering System that helps streamline restaurant operations. The system is designed to manage various tasks like menu management, order processing, payment handling, and kitchen operations. It utilizes **gRPC** for internal communication and exposes an efficient **GraphQL API layer** for interaction with the frontend.

### Key Features:
- **Menu Management**: Manage items, pricing, and availability.
- **User Authentication**: User login, registration, and account management.
- **Order Tracking**: Real-time order processing and status updates.
- **Payment Integration**: Secure payment processing for orders.
- **Kitchen Operations**: Order fulfillment based on priority and status.

## Architecture Diagram
![Architecture Diagram](https://github.com/CallenCaracy/ByteBites/blob/main/documents/diagrams/SIA_Final_Project.drawio.png)

## Tech Stack
- **Frontend**: React with Tailwind CSS (Vite as build tool)
- **Backend**: Go (GoLang)
- **Database**: PostgreSQL (Hosted on Supabase)
- **API Layer**: GraphQL (GQLGen)
- **Internal Communication**: gRPC
- **Database Migration**: Goose

---

## Microservices Architecture

### 1. **Menu Service**
   - **Responsibilities**: Manages the menu items and employee accounts.
   - **Storage**: Menu data, pricing, and availability are stored in the **Menu DB**.
   - **Tech**: GoLang for implementation, gRPC for communication.
   
### 2. **User Service**
   - **Responsibilities**: Handles user authentication and account management.
   - **Storage**: User data is stored in the **User DB**.
   - **Tech**: GoLang for implementation, gRPC for communication.
   
### 3. **Order Service**
   - **Responsibilities**: Handles order placement and tracking, integrates with the Payment Service.
   - **Storage**: Order data is stored in the **Order DB**.
   - **Tech**: GoLang for implementation, gRPC for communication.

### 4. **Payment Service**
   - **Responsibilities**: Processes payments securely.
   - **Storage**: Payment data is stored in the **Payment DB**.
   - **Tech**: GoLang for implementation, gRPC for communication.

### 5. **Kitchen Service**
   - **Responsibilities**: Manages order fulfillment and kitchen operations.
   - **Tech**: GoLang for implementation, gRPC for communication.

### 6. **GraphQL Service**
   - **Responsibilities**: Acts as the main API gateway for frontend communication.
   - **Tech**: GraphQL for data aggregation, GoLang for implementation, and gRPC for internal microservice communication.

---

## Installation & Setup

### Prerequisites
Before you begin, ensure you have the following tools installed:

- [Go](https://go.dev/dl/)
- [Node.js](https://nodejs.org/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Supabase](https://supabase.com/) for database hosting
- [Goose](https://github.com/pressly/goose) for database migrations

---

### Clone the Repository
```sh
git clone https://github.com/CallenCaracy/ByteBites.git
cd ByteBites
