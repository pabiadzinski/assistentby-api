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

func (api *ApiAssistent) StoreBankAccount(form BankAccount) interface{} {
	response := &ApiResponse{Data: BankAccount{}}

	req, err := api.C.postRequest(api.C.getEndpoint(storeBankAccount), form)

	failOnError(err, "create bank account")

	json.Unmarshal(req, response)

	return response
}
