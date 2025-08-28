package ledger

type CreateAccountRequest struct {
	AccountName string       `json:"account_name" binding:"required,min=3,max=255"`
	Currency    CurrencyType `json:"currency" binding:"required,oneof=SAR USD"`
}
