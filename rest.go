package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type method string

const (
	getOperations = method("operations")
	getUser       = method("user")
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type ClientApi struct {
	c         Doer // ClientApi
	endpoints map[method]string
	token     string
	baseUrl   string
	teamId    string
}

func NewClient(baseUrl string, teamId string, token string) *ClientApi {
	return &ClientApi{
		c:         &http.Client{},
		endpoints: createEndpoints(baseUrl),
		baseUrl:   baseUrl,
		teamId:    teamId,
		token:     token,
	}
}

func (c *ClientApi) validateResponse(res *http.Response) ([]byte, error) {
	ct := res.Header.Get(http.CanonicalHeaderKey("Content-Type"))

	if ct != "application/json" {
		return nil, errors.New("response is not json")
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *ClientApi) getEndpoint(method method) string {
	endpoint, ok := c.endpoints[method]

	if !ok {
		panic(fmt.Errorf("endpoint %s not found", string(method)))
	}

	return endpoint
}

func createEndpoints(baseURI string) map[method]string {
	list := map[method]string{}

	list[getOperations] = fmt.Sprint(baseURI, "/", string(getOperations))
	list[getUser] = fmt.Sprint(baseURI, "/", string(getUser))

	return list
}

func (c *ClientApi) postRequest(url string, postData interface{}) ([]byte, error) {
	j, _ := json.Marshal(postData)
	form := bytes.NewBuffer(j)

	req, err := http.NewRequest("POST", url, form)

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", c.token)

	response, err := c.c.Do(req)

	defer response.Body.Close()

	res, err := c.validateResponse(response)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *ClientApi) getRequest(url string, params string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", c.token)

	response, err := c.c.Do(req)

	defer response.Body.Close()

	res, err := c.validateResponse(response)

	if err != nil {
		return nil, err
	}

	return res, err
}
