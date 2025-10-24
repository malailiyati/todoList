# 🧩 Todo List Backend (Go + Gin + PostgreSQL)

A RESTful API backend for managing todos and categories, built using **Go (Gin Framework)** and **PostgreSQL**.

This backend powers the Todo List Frontend by providing endpoints for creating, updating, deleting, and filtering todos and categories.

## ⚙️ Tech Stack

| Category    | Technology                        |
| ----------- | --------------------------------- |
| Language    | Go (Golang)                       |
| Framework   | Gin                               |
| Database    | PostgreSQL                        |
| ORM         | GORM                              |
| Environment | godotenv                          |
| Build Tool  | Makefile (for migrations & setup) |

## 🧭 Project Structure

```
.
├── cmd/
│   └── main.go                  → entry point (server setup)
├── internals/
│   ├── configs/                 → database configuration
│   ├── handlers/                → route handlers (controllers)
│   ├── models/                  → data models (Todo, Category)
│   ├── repositories/            → database queries (GORM)
│   ├── routers/                 → Gin route registration
│   ├── services/                → business logic layer
│   └── utils/                   → response helpers (Success/Error)
├── migrations/                  → SQL schema or Makefile commands
├── go.mod / go.sum              → dependencies
└── .env                         → environment variables
```

## ⚡ Installation & Setup

### 1️⃣ Clone repository

```bash
git clone https://github.com/malailiyati/todoList.git
cd todoList
```

### 2️⃣ Set up environment variables

Buat file `.env`:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=todo_db
DB_SSLMODE=disable
```

### 3️⃣ Run PostgreSQL & create database

```bash
createdb todo_db
```

### 4️⃣ Run migrations (via Makefile)

```bash
make migrate-up
```

### 5️⃣ Run server

```bash
go run cmd/main.go
```

Server akan berjalan di:

```
http://localhost:8080
```

## 🗄️ Database Schema

### 🧱 Table: categories

| Column     | Type               | Description            |
| ---------- | ------------------ | ---------------------- |
| id         | SERIAL PRIMARY KEY | unique identifier      |
| name       | VARCHAR(100)       | category name (unique) |
| color      | VARCHAR(25)        | hex or named color     |
| created_at | TIMESTAMPTZ        | record creation time   |
| updated_at | TIMESTAMPTZ        | record update time     |

### 🧱 Table: todos

| Column      | Type               | Description                          |
| ----------- | ------------------ | ------------------------------------ |
| id          | SERIAL PRIMARY KEY | unique identifier                    |
| title       | VARCHAR(255)       | todo title                           |
| description | TEXT               | task description                     |
| completed   | BOOLEAN            | completion status                    |
| category_id | INT                | foreign key to categories (nullable) |
| priority    | VARCHAR(10)        | high / medium / low                  |
| due_date    | TIMESTAMPTZ        | optional deadline                    |
| created_at  | TIMESTAMPTZ        | record creation time                 |
| updated_at  | TIMESTAMPTZ        | record update time                   |

### 🪄 Relation:

**One Category → Many Todos**

(Todos with deleted category will have `category_id = NULL`)

## 🔌 API Endpoints

### 📂 Categories

| Method | Endpoint              | Description         |
| ------ | --------------------- | ------------------- |
| GET    | `/api/categories`     | Get all categories  |
| POST   | `/api/categories`     | Create new category |
| PATCH  | `/api/categories/:id` | Update category     |
| DELETE | `/api/categories/:id` | Delete category     |

### ✅ Todos

| Method | Endpoint                  | Description                          |
| ------ | ------------------------- | ------------------------------------ |
| GET    | `/api/todos`              | Get todos (with search + pagination) |
| GET    | `/api/todos/:id`          | Get todo by ID                       |
| POST   | `/api/todos`              | Create new todo                      |
| PATCH  | `/api/todos/:id`          | Update todo                          |
| PATCH  | `/api/todos/:id/complete` | Toggle todo completion               |
| DELETE | `/api/todos/:id`          | Delete todo                          |

## 🔍 Features Implemented

- **Search & Pagination** (`/api/todos?search=...&page=...`)
- **Priority Sorting**: high → medium → low
- **Completion Sorting**: unfinished tasks first
- **Category Relationship**: cascade updates, set null on delete
- **Validation Rules**:
  - **Category**: name unique, color valid (`#hex` or color name)
  - **Todo**: title required, valid priority & category_id

## 🧩 Code Architecture

| Layer                | Description                           |
| -------------------- | ------------------------------------- |
| Handler (Controller) | Handle API requests & responses       |
| Service              | Contain business logic and validation |
| Repository           | Database access via GORM              |
| Model                | Define data structure & relationships |
| Utils                | Unified JSON response helper          |
| Configs              | Database initialization using .env    |

**Pattern used**: Handler → Service → Repository → Database

## 🧪 Error Handling

No middleware used — each handler returns JSON using:

- `utils.Success()` → for success responses
- `utils.Error()` → for error responses

This keeps API responses consistent and easy to parse by the frontend.

## 🌱 Future Improvements

- Add unit tests for service & repository layers
- Add due date filtering
- Add sorting endpoints for flexible UI control
- Implement caching (e.g., Redis) for larger datasets

## 👩‍💻 Author

**Ma'la Iliyati**  
Full Stack Developer | Go • React • PostgreSQL • Docker

- 🔗 [GitHub](https://github.com/malailiyati)
- 🔗 [LinkedIn](https://www.linkedin.com/in/ma-la-iliyati)

**Related Project:**  
🔗 [Todo List Frontend (React + Ant Design)](https://github.com/malailiyati/todoList-fe.git)

## 📄 License

This project is licensed under the **MIT License**.
