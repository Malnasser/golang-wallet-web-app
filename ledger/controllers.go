package ledger

import (
	"net/http"
	database "simple/payment-wallet/core"

	"github.com/gin-gonic/gin"
)

func SetupRouter(rg *gin.RouterGroup) {
	ledgersGroup := rg.Group("/account")

	ledgersGroup.POST("/", createAccount)
}

func createAccount(c *gin.Context) {
	var req CreateAccountRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	account := Account{
		AccountName: req.AccountName,
		Currency:    req.Currency,
	}

	if err := database.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create account",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account created successfully",
		"account": account,
	})
}
