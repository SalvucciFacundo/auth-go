# Go Auth System - Hexagonal Architecture 🚀

A professional Authentication System built with **Go**, following **Hexagonal Architecture (Ports & Adapters)** principles to ensure a decoupled, testable, and maintainable codebase.

## 🏗️ Architecture
The project adheres to Hexagonal Architecture standards:
- **Domain**: Entities and Ports (interfaces). Zero external dependencies.
- **Application**: Services and Business Logic (Use cases).
- **Infrastructure**: Concrete implementations (Adapters) for HTTP and Persistence (PostgreSQL).

## ✨ Features
- [x] User Registration with secure password hashing (`bcrypt`).
- [x] **JWT** (JSON Web Tokens) based authentication.
- [x] Security Middlewares for protected route management.
- [x] Password management (Secure password change).
- [x] PostgreSQL persistence using `pgx/v5`.

## 🛠️ Tech Stack
- **Language**: Go 1.22+
- **Database**: PostgreSQL
- **Security**: JWT (v5), Bcrypt
- **DB Driver**: pgx (v5)

## 🚀 Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/SalvucciFacundo/auth-go.git
cd auth-go
```

### 2. Database Setup
Ensure PostgreSQL is running and execute the `schema.sql` script:
```bash
psql -U your_user -d auth_db -f schema.sql
```

### 3. Run the Application
```bash
go run cmd/api/main.go
```

## 🛣️ API Endpoints
- `POST /register`: Register new users.
- `POST /login`: Authenticate and receive a JWT token.
- `GET /me`: (Protected) Authenticated user profile.
- `POST /change-password`: Secure password update.

---
Built with ❤️ for idiomatic Go practice.
