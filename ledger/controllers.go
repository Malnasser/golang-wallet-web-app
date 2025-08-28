package ledger

import (
	"errors"
	"net/http"
	"simple/payment-wallet/core"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SetupRouter(rg *gin.RouterGroup) {
	ledgersGroup := rg.Group("/accounts")

	ledgersGroup.POST("/", createAccount)
	ledgersGroup.GET("/", listAccounts)
	ledgersGroup.POST("/:account_uuid/top-up", topUpAccount)
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
// @Router /accounts [post]
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

// ListAccounts godoc
// @Summary List all accounts with pagination
// @Description Get a paginated list of accounts
// @Tags Accounts
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)" default(1)
// @Param page_size query int false "Number of items per page (default: 10)" default(10)
// @Success 200 {object} ListAccountsResponse "Accounts retrieved successfully"
// @Failure 400 {object} ErrorResponse "Bad request - invalid pagination parameters"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /accounts [get]
func listAccounts(c *gin.Context) {
	var req core.PaginationRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrorResponse{
			Error: "Invalid request data",
		})
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	var accounts []Account
	var totalCount int64
	if err := core.DB.Model(&Account{}).Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, core.ErrorResponse{
			Error: "Internal Error",
		})
	}

	offset := (req.Page - 1) * req.PageSize
	totalPages := int((totalCount + int64(req.PageSize) - 1) / int64(req.PageSize))
	if err := core.DB.Offset(offset).Limit(req.PageSize).Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, core.ErrorResponse{
			Error: "Failed to retrieve accounts",
		})
	}

	response := ListAccountsResponse{
		Accounts: accounts,
		Pagination: core.PaginationResponse{
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalCount: totalCount,
			TotalPages: totalPages,
		},
	}

	c.JSON(http.StatusOK, response)
}

// TopUpAccount godoc
// @Summary Top up an account
// @Description Add funds to a specific account by account ID
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account_uuid path string true "Account UUID" example(cff6f6f3-e440-460f-99e5-7ee0d4f20125)
// @Param topup body TopUpAccountRequest true "Top up details"
// @Success 200 {object} TopUpAccountResponse "Account topped up successfully"
// @Failure 400 {object} ErrorResponse "Bad request - invalid input"
// @Failure 404 {object} ErrorResponse "Account not found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /accounts/{account_uuid}/top-up [post]
func topUpAccount(c *gin.Context) {
	accountIDStr := c.Param("account_uuid")
	accountID, err := uuid.Parse(accountIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.ErrorResponse{
			Error: "Invalid account_uuid format",
		})
		return
	}

	var req TopUpAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrorResponse{
			Error: "Invalid request data",
		})
	}

	tx := core.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingTransaction Transaction
	err = tx.Where("idempotency_id = ?", req.IdempotencyID).First(&existingTransaction).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, core.ErrorResponse{
			Error: "Error checking idempotency",
		})
		return
	}
	if err == nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, core.ErrorResponse{
			Error: "Can't process duplicate transaction",
		})
	}

	var account Account
	if err := tx.Where("uuid = ?", accountID).First(&account).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, core.ErrorResponse{
			Error: "Account Not Found",
		})
	}

	account.Balance += req.Amount
	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, core.ErrorResponse{
			Error: "Failed to update account balance",
		})
	}

	transaction := Transaction{
		AccountUUID:   accountID,
		Amount:        req.Amount,
		AfterBalance:  account.Balance,
		TrxType:       TOP_UP,
		IdempotencyID: req.IdempotencyID,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, core.ErrorResponse{
			Error: "Failed to commit transaction",
		})
		return
	}
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusOK, TopUpAccountResponse{
			Message:     "Account topped up successfully",
			Account:     account,
			Transaction: transaction,
		})
	}
}
