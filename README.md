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

---

## Step 2: Create the Database

Create a file named `database.go`

```go
package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./pos.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create tables
	createTables := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		price REAL,
		stock INTEGER
	);

	CREATE TABLE IF NOT EXISTS sales (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_id INTEGER,
		quantity INTEGER,
		total REAL,
		sold_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(createTables)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
```

## Step 3: Set Up the Web Server

Create `main.go`:

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	db := InitDB()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "POS Dashboard API is running")
	})

	// Routes
	RegisterProductRoutes(r, db)
	RegisterSalesRoutes(r, db)

	r.Run(":8080")
}
```

## Step 4: Product API (CRUD)

Create `products.go`:

```go
package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterProductRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/products", func(c *gin.Context) {
		rows, _ := db.Query("SELECT id, name, price, stock FROM products")
		var products []map[string]interface{}

		for rows.Next() {
			var id int
			var name string
			var price float64
			var stock int
			rows.Scan(&id, &name, &price, &stock)
			products = append(products, gin.H{"id": id, "name": name, "price": price, "stock": stock})
		}

		c.JSON(http.StatusOK, products)
	})

	r.POST("/products", func(c *gin.Context) {
		var p struct {
			Name  string  `json:"name"`
			Price float64 `json:"price"`
			Stock int     `json:"stock"`
		}
		c.BindJSON(&p)
		stmt, _ := db.Prepare("INSERT INTO products(name, price, stock) VALUES (?, ?, ?)")
		_, err := stmt.Exec(p.Name, p.Price, p.Stock)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
	})
}
```