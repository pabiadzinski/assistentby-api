package assistentby

import (
	"encoding/json"
)

const BaseApiUrl = "https://app.assistent.by/api/v1"

type ApiAssistent struct {
	C ClientApi
}

func (api *ApiAssistent) GetOperations(params string) interface{} {
	response := &ApiResponse{Data: []Operation{}}

	req, err := api.C.getRequest(api.C.getEndpoint(getOperations), params)

	failOnError(err, "get operations")

	json.Unmarshal(req, response)

	//fmt.Println(reflect.TypeOf(response))
	//fmt.Printf("%+v\n", response)

	return response
}

func (api *ApiAssistent) GetProfile(params string) interface{} {
	response := &Profile{}

	req, err := api.C.getRequest(api.C.getEndpoint(getProfile), params)

	failOnError(err, "get user")

	json.Unmarshal(req, response)

	return response
}

func (api *ApiAssistent) StoreOperation(form Operation) interface{} {
	response := &ApiResponse{Data: Operation{}}

	req, err := api.C.postRequest(api.C.getEndpoint(storeOperation), form)

	failOnError(err, "create operation")

	json.Unmarshal(req, response)

	return response
}

func (api *ApiAssistent) UpdateOperation(form Operation, id string) interface{} {
	response := &ApiResponse{Data: Operation{}}

	req, err := api.C.putRequest(replaceId(api.C.getEndpoint(updateOperation), id), form)

	failOnError(err, "update operation")

	json.Unmarshal(req, response)

	return response
}

func (api *ApiAssistent) StoreBankAccount(form BankAccount) *BankAccountResponse {
	response := &BankAccountResponse{}

	req, err := api.C.postRequest(api.C.getEndpoint(storeBankAccount), form)

	failOnError(err, "create bank account")

	json.Unmarshal(req, response)

	return response
}

func (api *ApiAssistent) UpdateBankAccount(form BankAccount, id string) *BankAccountResponse {
	response := &BankAccountResponse{}

	req, err := api.C.putRequest(replaceId(api.C.getEndpoint(updateBankAccount), id), form)

	failOnError(err, "update bank account")

	json.Unmarshal(req, response)

	return response
}

func (api *ApiAssistent) StoreContractor(form Contractor) interface{} {
	response := &ApiResponse{Data: Contractor{}}

	req, err := api.C.postRequest(api.C.getEndpoint(storeContractor), form)

	failOnError(err, "create contractor")

	json.Unmarshal(req, response)

	return response
}

func (api *ApiAssistent) GetCurrencies(params string) *CurrencyResponse {
	response := &CurrencyResponse{}

	req, err := api.C.getRequest(api.C.getEndpoint(getCurrencies), params)

	failOnError(err, "get currencies")

	json.Unmarshal(req, response)

	return response
}

func (api *ApiAssistent) GetBanks(params string) *BankResponse {
	response := &BankResponse{}

	req, err := api.C.getRequest(api.C.getEndpoint(getBanks), params)

	failOnError(err, "get banks")

	json.Unmarshal(req, response)

	return response
}
