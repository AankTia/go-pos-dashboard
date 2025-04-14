package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterSalesRouter(r *gin.Engine, db *sql.DB) {
	r.POST("/sales", func(c *gin.Context) {
		var sale struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		}
		c.BindJSON(&sale)

		var price float64
		db.QueryRow("SELECT price FROM products WHERE id = ?", sale.ProductID).Scan(&price)
		total := price * float64(sale.Quantity)

		_, err := db.Exec("INSERT INTO sales(product_id, quantity, total) VALUES (?,?,?)", sale.ProductID, sale.Quantity, total)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		db.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", sale.Quantity, sale.ProductID)

		c.JSON(http.StatusOK, gin.H{"message": "Sales recorder", "total": total})
	})
}
