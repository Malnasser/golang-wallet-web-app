package ledger

import (
	"net/http"
	"simple/payment-wallet/core"

	"github.com/gin-gonic/gin"
)

func SetupRouter(rg *gin.RouterGroup) {
	ledgersGroup := rg.Group("/account")

	ledgersGroup.POST("/", createAccount)
}

// @BasePath /api/v1

// CreateAccount godoc
// @Summary Create a new account
// @Description Create a new account (payment wallet) for the user
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account body CreateAccountRequest true "Account creation details"
// @Success 201 {object} CreateAccountResponse "Account created successfully"
// @Failure 400 {object} ErrorResponse "Bad request - invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /account [post]
func createAccount(c *gin.Context) {
	var req CreateAccountRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrorResponse{
			Error:   "Invalid request data",
			Message: err.Error(),
		})
		return
	}

	account := Account{
		AccountName: req.AccountName,
		Currency:    req.Currency,
	}

	if err := core.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, core.ErrorResponse{
			Error: "Failed to create account",
		})
		return
	}

	c.JSON(http.StatusCreated, CreateAccountResponse{
		Message: "Account created successfully",
		Account: account,
	})
}
