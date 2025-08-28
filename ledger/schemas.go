package ledger

import "simple/payment-wallet/core"

type CreateAccountRequest struct {
	AccountName string       `json:"account_name" binding:"required,min=3,max=255"`
	Currency    CurrencyType `json:"currency" binding:"required,oneof=SAR USD"`
} //@name CreateAccountRequest

type CreateAccountResponse struct {
	Message string  `json:"message" example:"Account created successfully"`
	Account Account `json:"account"`
} //@name CreateAccountResponse

type ListAccountsResponse struct {
	Accounts   []Account               `json:"accounts"`
	Pagination core.PaginationResponse `json:"pagination"`
} //@name ListAccountsResponse

type TopUpAccountRequest struct {
	Amount        int64  `json:"amount" example:"42"`
	IdempotencyID string `json:"idempotencyId" example:"431e0419-23f3-41e4-a27e-a89815948be1"`
} // @name TopUpAccountRequest

type TopUpAccountResponse struct {
	Message     string      `json:"message" example:"Account topped up successfully"`
	Account     Account     `json:"Account"`
	Transaction Transaction `json:"Transaction"`
} // @name TopUpAccountResponse
