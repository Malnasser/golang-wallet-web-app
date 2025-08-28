package main

import (
	"log"
	"simple/payment-wallet/core"
	"simple/payment-wallet/ledger"

	"github.com/gin-gonic/gin"
)

func main() {
	config := core.LoadConfig()
	core.ConnectToDB(config)

	r := gin.Default()

	// config is now part of fin ctx
	r.Use(func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	})

	v1 := r.Group("/api/v1")
	ledger.SetupRouter(v1)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
