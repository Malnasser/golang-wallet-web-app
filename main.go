package main

import (
	"log"
	database "simple/payment-wallet/core"
	"simple/payment-wallet/ledger"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectToDB()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	ledger.SetupRouter(v1)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
