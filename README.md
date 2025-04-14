# Web-based Point of Sale (POS) dashboard

Here’s a step-by-step for building a web-based Point of Sale (POS) dashboard using Go (Golang) and SQLite. This setup will include basic product management, sales recording, and a simple dashboard UI.

---

## Prerequisites

- Go installed (1.20+)
- SQLite3 installed
- Basic knowledge of HTML, CSS, and JavaScript (vanilla or minimal React)
- Postman or curl for testing APIs
- Gin (Go web framework)

---

## Step 1: Project Setup

```bash
mkdir go-pos-dashboard
cd go-pos-dashboard
go mod init go-pos-dashboard
```

Install dependencies:

```bash
go get github.com/gin-gonic/gin
go get github.com/mattn/go-sqlite3
```