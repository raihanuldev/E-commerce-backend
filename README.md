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
## 📘 API Endpoints

Below are the available API endpoints for the **E-Commerce (GoLang)** project.  
Each endpoint includes the **method**, **URL**, and **example request body** (if applicable).

---
## 📘 API Endpoints

Below are the available API endpoints for the **E-Commerce (GoLang)** project.
Each endpoint includes the **method**, **URL**, and **example request body** (if applicable).

---

### 🧍 Auth

#### 1. Create User

**POST** `/users`
**Body:**

```json
{
  "frist_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "password": "12345",
  "is_shop_owner": true
}
```

#### 2. Login

**POST** `/login`
**Body:**

```json
{
  "email": "john@example.com",
  "password": "12345"
}
```

---

### 🛒 Cart

#### 1. Add to Cart

**POST** `/cart?userid=11`
**Body:**

```json
{
  "product_id": 133,
  "price": 700,
  "quantity": 22
}
```

---

### 📦 Product

#### 1. Get All Products

**GET** `/products?page=6&limit=10`

#### 2. Get Product by ID

**GET** `/products/{id}`
Example: `/products/1`

#### 3. Create Product

**POST** `/products`
**Headers:**

```
Authorization: Bearer <token>
```

**Body:**

```json
{
  "title": "Kala pahki",
  "description": "This is Best Quality Headphone I have ever seen",
  "price": 167.25,
  "ImgUrl": "https://image.url"
}
```

---

### 📦 Orders

#### 1. Create Order

**POST** `/order`
**Body:**

```json
{
  "productId": 12,
  "userId": 7,
  "quantity": 3,
  "status": "pending"
}
```

#### 2. Get All Orders

**GET** `/order?page=6&limit=10`

#### 3. Update Order Status

**PUT** `/order`
**Body:**

```json
{
  "orderId": 2,
  "status": "in-progress"
}
```
