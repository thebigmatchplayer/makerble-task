# 🏥 Markerble Portal — Role-based Patient Management API

A RESTful backend service built in **Go** for managing patients in a hospital setup. Supports **JWT-based authentication**, **role-based access control (RBAC)**, and follows clean architecture patterns.

---

## 🔧 Tech Stack

- **Language:** Golang
- **Database:** PostgreSQL (Neon DB used)
- **Auth:** JWT (JSON Web Token)
- **Access Control:** YAML-based Access Control Matrix (ABAC-style)
- **Routing:** net/http
- **Validation:** go-playground/validator
- **Logging:** zap (Uber-go)
---

## 🚦 Control Flow

1. **User Registration/Login**
   - `POST /register` creates a user with hashed password (bcrypt).
   - `POST /login` returns a JWT token after validation.

2. **Authentication Middleware**
   - Extracts and validates JWT from the `Authorization` header.
   - Decodes claims (user ID, role) and performs role-based checks.

3. **Authorization**
   - Controlled via `access.yaml`, where each `METHOD` + `PATH` maps to allowed roles.
   - E.g.:
     ```yaml
     /patients:
       POST: ["receptionist"]
       GET: ["doctor", "receptionist"]
     ```

4. **Route Handling**
   - All routes are set up via `handlers.SetupRoutes()` using `net/http`.
   - A single endpoint (`/patients`) handles GET, POST, PUT, DELETE with role checks.

---

## ✅ Features

- 🔐 **JWT-based Auth** with expiry and claims (role, user ID)
- ✅ **Role-based Access** for `doctor` and `receptionist`
- 🗄️ **PostgreSQL Integration** using standard `database/sql`
- ⚙️ **Middleware-based Logging** using `zap`
- 📋 **Validation** for JSON inputs with `validator.v10`

---

## 🚀 Endpoints

| Endpoint        | Method | Roles Allowed     | Description                       |
|----------------|--------|-------------------|-----------------------------------|
| /register       | POST   | Public            | Register a new user               |
| /login          | POST   | Doctor, Receptionist            | Login and receive JWT             |
| /patients       | GET    | Doctor, Receptionist | View all or specific patients     |
| /patients       | POST   | Receptionist      | Create new patient                |
| /patients       | PUT    | Doctor            | Update patient diagnosis          |
| /patients       | DELETE | Receptionist      | Delete patient                    |
---
## 🧠 Best Practices Followed

- ✅ **Layered architecture**: Separation of models, handlers, config, and middleware
- ✅ **Access Control Matrix (ACM)**: YAML-based declarative permission mapping
- ✅ **JWT**: Short-lived tokens with embedded claims (user ID, role)
- ✅ **Hashing passwords**: Using bcrypt before storing in DB
- ✅ **Zap logging**: For debug/info/warning/error logs across all layers
- ✅ **Environment variables**: DB credentials kept secure via `.env`
---
## 📦 Packages Used

### 🔹 Standard Library

- [`net/http`](https://pkg.go.dev/net/http) – HTTP server and routing
- [`os`](https://pkg.go.dev/os) – Environment variable loading and file operations
- [`strings`](https://pkg.go.dev/strings) – String utilities (e.g. trimming, casing)
- [`database/sql`](https://pkg.go.dev/database/sql) – Core SQL DB interaction
- [`encoding/json`](https://pkg.go.dev/encoding/json) – JSON parsing and encoding
- [`errors`](https://pkg.go.dev/errors) – Error wrapping and comparison
- [`fmt`](https://pkg.go.dev/fmt) – String formatting

### 🔹 Third-Party

- [`github.com/go-playground/validator`](https://github.com/go-playground/validator) – Struct validation for incoming JSON payloads
- [`github.com/golang-jwt/jwt/v5`](https://github.com/golang-jwt/jwt) – JWT generation and parsing
- [`golang.org/x/crypto/bcrypt`](https://pkg.go.dev/golang.org/x/crypto/bcrypt) – Password hashing
- [`go.uber.org/zap`](https://github.com/uber-go/zap) – Structured, high-performance logging
- [`github.com/joho/godotenv`](https://github.com/joho/godotenv) – Load environment variables from `.env` file
- [`github.com/lib/pq`](https://github.com/lib/pq) – PostgreSQL driver for `database/sql`
- [`gopkg.in/yaml.v3`](https://pkg.go.dev/gopkg.in/yaml.v3) – Parse and manage the Access Control Matrix in YAML
