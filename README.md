# Go CRUD API

A simple Go-based CRUD project using Gin, GORM, and PostgreSQL.

## 🛠 Setup Instructions

### 1. Initialize Go Module

```bash
go mod init github.com/mcharolabs/go-crud
```

### 2. Install Development Tools

These tools are used during development (they won't appear in go.mod unless manually added via tools.go).
Install CompileDaemon (live reload)

```bash
go install github.com/githubnemo/CompileDaemon@latest
```

Install godotenv (for .env loading via CLI)

```bash
go install github.com/joho/godotenv/cmd/godotenv@latest
```

Ensure your $GOPATH/bin or $GOBIN is in your $PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### 3. Install Project Dependencies

Gin (web framework)

```bash
go get github.com/gin-gonic/gin
```

GORM (ORM library)

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### 4. Running the App with Auto Reload

Use `CompileDaemon` to auto-restart the server on code changes:

```bash
CompileDaemon -command="./go-crud"
```
