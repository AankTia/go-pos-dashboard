package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterProductRouter(r *gin.Engine, db *sql.DB) {
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

		stmt, _ := db.Prepare("INSERT INTO products(name, proce, stock) VALUES (?,?,?)")
		_, err := stmt.Exec(p.Name, p.Price, p.Stock)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"massage": "Product created"})
	})
}
