package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := InitDB()

	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "POS Dashboard API is running")
	// })

	r.GET("/dashboard", func(c *gin.Context) {
		row := db.QueryRow("SELECT COUNT(*), SUM(total) FROM sales")
		var count int
		var total float64
		row.Scan(&count, &total)

		c.JSON(http.StatusOK, gin.H{
			"total_sales": count,
			"revenue":     total,
		})
	})

	// Routes
	RegisterProductRouter(r, db)
	RegisterSalesRouter(r, db)

	r.StaticFile("/", "./index.html")

	r.Run(":8080")
}
