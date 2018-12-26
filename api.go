package assistentby

import (
	"encoding/json"
	_ "fmt"
)

const BaseApiUrl = "https://app.assistent.by/api/v1"

type ApiAssistent struct {
	C ClientApi
}

func (api *ApiAssistent) GetOperations(params string) (*ApiResponse, error) {
	response := &ApiResponse{Data: []Operation{}}

	req, err := api.C.getRequest(api.C.getEndpoint(getOperations), params)

	failOnError(err, "get operations")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) GetProfile(params string) (*Profile, error) {
	response := &Profile{}

	req, err := api.C.getRequest(api.C.getEndpoint(getProfile), params)

	failOnError(err, "get user")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) StoreOperation(form Operation) (*ApiResponse, error) {
	response := &ApiResponse{Data: Operation{}}

	req, err := api.C.postRequest(api.C.getEndpoint(storeOperation), form)

	failOnError(err, "create operation")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) UpdateOperation(form Operation, id string) (*ApiResponse, error) {
	response := &ApiResponse{Data: Operation{}}

	req, err := api.C.putRequest(replaceId(api.C.getEndpoint(updateOperation), id), form)

	failOnError(err, "update operation")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) StoreBankAccount(form BankAccount) (*BankAccountResponse, error) {
	response := &BankAccountResponse{}

	req, err := api.C.postRequest(api.C.getEndpoint(storeBankAccount), form)

	failOnError(err, "create bank account")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) UpdateBankAccount(form BankAccount, id string) (*BankAccountResponse, error) {
	response := &BankAccountResponse{}

	req, err := api.C.putRequest(replaceId(api.C.getEndpoint(updateBankAccount), id), form)

	failOnError(err, "update bank account")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) StoreContractor(form Contractor) (*ApiResponse, error) {
	response := &ApiResponse{Data: Contractor{}}

	req, err := api.C.postRequest(api.C.getEndpoint(storeContractor), form)

	failOnError(err, "create contractor")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) GetCurrencies(params string) (*CurrencyResponse, error) {
	response := &CurrencyResponse{}

	req, err := api.C.getRequest(api.C.getEndpoint(getCurrencies), params)

	failOnError(err, "get currencies")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) GetBanks(params string) (*BankResponse, error) {
	response := &BankResponse{}

	req, err := api.C.getRequest(api.C.getEndpoint(getBanks), params)

	failOnError(err, "get banks")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) UpdateTransactionStatus(form TransactionStatus) (*PaymentStatusResponse, error) {
	response := &PaymentStatusResponse{}

	req, err := api.C.postRequest(api.C.getEndpoint(updateTransactionStatus), form)

	failOnError(err, "update transaction status")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) UpdateAccountBalance(form BankAccountBalance) (*AccountBalanceResponse, error) {
	response := &AccountBalanceResponse{}

	req, err := api.C.postRequest(api.C.getEndpoint(updateBankAccountBalance), form)

	failOnError(err, "update account balance")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) StorePayment(form Payment) (*PaymentResponse, error) {
	response := &PaymentResponse{}

	req, err := api.C.postRequest(api.C.getEndpoint(storePayment), form)

	failOnError(err, "create payment")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) UpdateTeam(form Team, id string) (*TeamResponse, error) {
	response := &TeamResponse{}

	req, err := api.C.putRequest(replaceId(api.C.getEndpoint(updateTeam), id), form)
	failOnError(err, "update team")

	json.Unmarshal(req, response)

	return response, err
}

func (api *ApiAssistent) StoreTransaction(form Operation) (*ApiResponse, error) {
	response := &ApiResponse{Data: Operation{}}

	req, err := api.C.postRequest(api.C.getEndpoint(storeTransaction), form)

	failOnError(err, "Create transaction")

	json.Unmarshal(req, response)

	return response, err
}
