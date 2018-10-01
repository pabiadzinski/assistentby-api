package assistentby

import (
	"encoding/json"
)

const BaseApiUrl = "https://app.assistent.by/api/v1"

type ApiAssistent struct {
	c ClientApi
}

func (api *ApiAssistent) getOperations(params string) interface{} {
	response := &ApiResponse{Data: []Operation{}}

	req, err := api.c.getRequest(api.c.getEndpoint(getOperations), params)

	failOnError(err, "get operations")

	json.Unmarshal(req, response)

	//fmt.Println(reflect.TypeOf(response))
	//fmt.Printf("%+v\n", response)

	return response
}

func (api *ApiAssistent) getProfile(params string) interface{} {
	response := &Profile{}

	req, err := api.c.getRequest(api.c.getEndpoint(getProfile), params)

	failOnError(err, "get user")

	json.Unmarshal(req, response)

	return response
}

func (api *ApiAssistent) addOperation(form Operation) interface{} {
	response := &ApiResponse{}

	req, err := api.c.postRequest(api.c.getEndpoint(storeOperation), form)

	failOnError(err, "create operation")

	json.Unmarshal(req, response)

	return response
}
