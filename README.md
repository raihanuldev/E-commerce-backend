# 🛒 E-Commerce Backend

A **Domain Driven Design (DDD)** based backend service for an E-Commerce application, built with **Golang** and **PostgreSQL**.  
The system implements clean architecture principles with repository pattern, modular design, and secure JWT authentication.

---

## ⚙️ Tech Stack

- **Language:** Go `1.25.0`
- **Framework/Architecture:** Domain Driven Design (DDD)
- **Database:** PostgreSQL
- **ORM / SQL Helper:** `sqlx`
- **Authentication:** JWT (JSON Web Token)
- **Environment Management:** `godotenv`
- **Dependency Management:** Go Modules

---

## 🧩 Features

- 🧍‍♂️ **User Management**
  - Create, Read, Update, Delete (CRUD)
  - Secure password hashing & JWT authentication

- 🛍 **Product Management**
  - Full CRUD operations for products
  - Category-based organization

- 🔐 **Authentication**
  - JWT token-based login & route protection
  - Token validation middleware

- 🗄 **Repository Pattern**
  - Decoupled data access layer for flexibility and testability

- 🧠 **Domain Driven Design**
  - Clear separation between domain, repository, and handler layers
  - Each domain (e.g. `user`, `product`) has its own model and service

---
## ⚙️ Concurrency Handling
The application leverages **Go’s goroutines** to run multiple operations concurrently — improving performance and responsiveness.  
Every concurrent routine is controlled using **sync.WaitGroup**, ensuring that all goroutines complete safely before the main process exits.  

This approach helps the system efficiently handle:
- Multiple user requests at the same time  
- Parallel data fetching or processing tasks  
- Resource-intensive background operations  

---

