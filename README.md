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
- **API TESTING:** Postman : https://rihanulislam2015-5709015.postman.co/workspace/Raihanul-Islam-Sharif's-Workspa~0cac0da8-57ea-456f-8ca1-99e74b11def3/collection/49116042-480e509a-b618-4974-a725-1abe8ee09385?action=share&creator=49116042&active-environment=49116042-4bf64501-6b73-4866-845f-0e659a4c8c55

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

