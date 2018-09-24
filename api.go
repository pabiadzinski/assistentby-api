package main

import (
	"encoding/json"
	"fmt"
	"github.com/pabiadzinski/assistentby-api/models"
	"reflect"
)

const BaseApiUrl = "https://app.assistent.by/api/v1"

type ApiAssistent struct {
	c *ClientApi
}

func (api *ApiAssistent) getOperations(params string) interface{} {
	response := &models.ApiResponse{Data: []models.Operation{}}

	url := api.c.baseUrl + "/operations?team_id=" + api.c.teamId

	req, err := api.c.getRequest(url, params)

	failOnError(err, "get operations")

	json.Unmarshal(req, response)

	fmt.Println(reflect.TypeOf(response))
	//fmt.Printf("%+v\n", response)

	return response
}

func (api *ApiAssistent) getUser(params string) interface{} {
	response := &models.Profile{}

	url := api.c.baseUrl + "/user"

	req, err := api.c.getRequest(url, params)

	failOnError(err, "get user")

	json.Unmarshal(req, response)

	return response
}

func (api *ApiAssistent) addOperation(form models.Operation) interface{} {
	response := &models.ApiResponse{Data: models.Operation{}}

	url := api.c.baseUrl + "/operations"

	req, err := api.c.postRequest(url, form)

	failOnError(err, "create operation")

	json.Unmarshal(req, response)

	return response
}
