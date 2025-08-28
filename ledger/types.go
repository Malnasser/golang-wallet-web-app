package ledger

const (
	SAR = "SAR"
	USD = "USD"
)

// Transaction type constants
const (
	CREDIT = "CREDIT"
	DEBIT  = "DEBIT"
	TOP_UP = "TOP_UP"
)

// Currency type for validation
type CurrencyType string

const (
	CurrencySAR CurrencyType = SAR
	CurrencyUSD CurrencyType = USD
)

// TransactionType for validation
type TransactionType string

const (
	TransactionCredit TransactionType = CREDIT
	TransactionDebit  TransactionType = DEBIT
	TransactionTopUp  TransactionType = TOP_UP
)
