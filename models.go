package assistentby

const (
	OperationIncomeID = iota + 1
	OperationExpenseID
)

const (
	OperationFlowCashID = iota + 1
	OperationFlowBankID
)

const (
	PaymentStatusCreated     = "created"
	PaymentStatusUnconfirmed = "unconfirmed"
	PaymentStatusConfirmed   = "confirmed"
	PaymentStatusConducted   = "conducted"
	PaymentStatusRejected    = "rejected"
	PaymentStatusProcessing  = "processing"
	PaymentStatusPostponed   = "postponed"
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

type BankResponse struct {
	Data         []Bank `json:"data"`
	MetaResponse `json:"meta"`
}

type CurrencyResponse struct {
	Data         []Currency `json:"data"`
	MetaResponse `json:"meta"`
}

type BankAccountResponse struct {
	Data         BankAccount `json:"data"`
	MetaResponse `json:"meta"`
}

type PaymentStatusResponse struct {
	Data         TransactionStatus `json:"data"`
	MetaResponse `json:"meta"`
}

type PaymentResponse struct {
	Data         Payment `json:"data"`
	MetaResponse `json:"meta"`
}

type AccountBalanceResponse struct {
	Data         BankAccount `json:"data"`
	MetaResponse `json:"meta"`
}

type TeamResponse struct {
	Data Team `json:"data"`
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
	UIN  string `json:"uin"`
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
	BankAccountId   string  `json:"bank_account_id"`
	ContractorId    string  `json:"contractor_id"`
	CurrencyId      string  `json:"currency_id"`
	TransactionId   string  `json:"transaction_id,omitempty"`
	CreatedBy       *string `json:"created_by,omitempty"`
	CreatedAt       *string `json:"created_at,omitempty"`
	UpdatedAt       *string `json:"updated_at,omitempty"`
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

type Payment struct {
	Id              string  `json:"id,omitempty"`
	Date            string  `json:"date"`
	Number          string  `json:"number"`
	TeamId          string  `json:"team_id,omitempty"`
	BankId          int     `json:"bank_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	PaymentStatusId int     `json:"payment_status_id,omitempty"`
	CurrencyId      int     `json:"currency_id,omitempty"`
	Subject         string  `json:"subject"`
	PaymentCode     string  `json:"payment_code"`
}

type TransactionStatus struct {
	TransactionId string `json:"transaction_id,omitempty"`
	TeamId        string `json:"team_id,omitempty"`
	Status        string `json:"status,omitempty"`
}

type BankAccountBalance struct {
	Number  string `json:"number,omitempty"`
	TeamId  string `json:"team_id,omitempty"`
	Balance int    `json:"balance,omitempty"`
}

type Contractor struct {
	Id            string `json:"id,omitempty"`
	TeamId        string `json:"team_id,omitempty"`
	Name          string `json:"name,omitempty"`
	ShortName     string `json:"short_name,omitempty"`
	UIN           string `json:"uin,omitempty"`
	IsLegalPerson int    `json:"is_legal_person,omitempty"`
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

type SalaryEmployee struct {
	BankAccount string `json:"bank_account"`
}

type Employee struct {
	Id             string `json:"id,omitempty"`
	TeamId         string `json:"team_id"`
	FistName       string `json:"fist_name"`
	MiddleName     string `json:"middle_name"`
	LastName       string `json:"last_name"`
	PassportNumber string `json:"passport_number"`
	PassportSeries string `json:"passport_series"`
	PersonalCode   string `json:"personal_code"`
	Position       string `json:"position"`
	Salary         string `json:"salary"`
	TariffRate     string `json:"tariff_rate"`
	BankAccount    string `json:"bank_account"`
}
