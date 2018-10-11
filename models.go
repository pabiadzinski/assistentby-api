package assistentby

const (
	OperationIncomeID = iota + 1
	OperationExpenseID
)

const (
	OperationFlowBankID = iota + 1
	OperationFlowCashID
)

type MetaResponse struct {
	ApiVersion         string `json:"api_version"`
	PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
	Total       int         `json:"total"`
	Count       int         `json:"count"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	Links       interface{} `json:"links"`
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type ApiResponse struct {
	Data         interface{} `json:"data"`
	MetaResponse `json:"meta"`
}

type ErrorResponse struct {
	Error `json:"error"`
}

type ApiErrorResponse struct {
	Error `json:"error"`
}

type Profile struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	OwnedTeams interface{} `json:"owned_teams"`
	Teams      interface{} `json:"teams"`
	TaxRate    int         `json:"tax_rate"`
}

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Operation struct {
	Id              string  `json:"id"`
	TeamId          string  `json:"team_id,omitempty"`
	Date            string  `json:"date"`
	AmountIn        float32 `json:"amount_in,omitempty"`
	AmountOut       float32 `json:"amount_out,omitempty"`
	PaymentDocument string  `json:"payment_document"`
	Description     string  `json:"description,omitempty"`
	OperationTypeId int     `json:"operation_type_id"`
	OperationFlowId int     `json:"operation_flow_id,omitempty"`
	CurrencyId      int     `json:"currency_id"`
	TransactionId   string  `json:"transaction_id,omitempty"`
	CreatedBy       string  `json:"created_by,omitempty"`
	CreatedAt       string  `json:"created_at,omitempty"`
	UpdatedAt       string  `json:"updated_at,omitempty"`
}

type BankAccount struct {
	Id                 string  `json:"id,omitempty"`
	TeamId             string  `json:"team_id,omitempty"`
	BankId             int     `json:"bank_id,omitempty"`
	Label              string  `json:"label,omitempty"`
	Balance            float64 `json:"balance,omitempty"`
	BalanceDate        string  `json:"balance_date,omitempty"`
	CurrencyId         int     `json:"currency_id,omitempty"`
	Number             string  `json:"number,omitempty"`
	OpeningBalance     float64 `json:"opening_balance,omitempty"`
	OpeningBalanceDate string  `json:"opening_balance_date,omitempty"`
}

type OperationType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Bank struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Status  string `json:"status"`
}

type Okved struct {
	Id    int    `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Currency struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	CodeName string `json:"code_name"`
}
