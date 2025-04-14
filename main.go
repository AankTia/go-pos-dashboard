package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db:= InitDB()

	r.GET("/", func (c *gin.Context)  {
		c.String(http.StatusOK, "POS Dashboard API is running")
	})

	// Routes
	RegisterProductRouter(r, db)
	RegisterSalesRouter(r, db)

	r.Run(":80880")
}