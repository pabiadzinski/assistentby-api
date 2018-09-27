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
	getProfile method = "user"

	getBankAccounts    method = "bank-accounts?team_id="
	storeBankAccounts  method = "bank-accounts"
	updateBankAccounts method = "bank-accounts/{id}"
	deleteBankAccounts method = "bank-accounts/{id}"

	getBusinesses  method = "businesses?team_id="
	storeBusiness  method = "businesses"
	updateBusiness method = "businesses/{id}"
	deleteBusiness method = "businesses/{id}"

	getContractors   method = "contractors?team_id="
	storeContractor  method = "contractors"
	updateContractor method = "contractors/{id}"
	deleteContractor method = "contractors/{id}"

	getCurrencies method = "currencies"

	getDocuments     method = "documents?team_id="
	getDocumentQueue method = "documents/{id}/queue"
	storeDocumentXml method = "documents/{id}/xml"
	sendDocument     method = "documents/{id}/send"
	approveDocument  method = "documents/{id}/approve"
	confirmDocument  method = "documents/{id}/confirm"
	shareDocument    method = "documents/{id}/share"
	storeDocument    method = "documents"
	updateDocument   method = "documents/{id}"
	deleteDocument   method = "documents/{id}"

	getDocumentStatues    method = "documents/statuses"
	getDocumentTypes      method = "documents/types?team_id="
	getDocumentOperations method = "documents/operations?team_id="
	getDocumentBills      method = "documents/bills?team_id="
	getOkveds             method = "okveds"
	getOkvedCategories    method = "okveds/categories"

	getOperations      method = "operations?team_id="
	storeOperation     method = "operations"
	updateOperation    method = "operations/{id}"
	deleteOperation    method = "operations/{id}"
	getLatestOperation method = "operations/latest"
	getOperationFlows  method = "operations/flows"
	getOperationTypes  method = "operations/types"
	getOperationStats  method = "operations/stats"

	storePushTokens  method = "push-tokens"
	deletePushTokens method = "push-tokens/{id}"

	getPayments        method = "payments?team_id="
	storePayment       method = "payments"
	storePaymentWord   method = "payments"
	updatePayment      method = "payments/{id}"
	deletePayment      method = "payments/{id}"
	getLatestPayment   method = "payments/latest"
	getPaymentStatuses method = "payments/statues"

	getServiceCategories method = "service-categories?team_id="
	getServices          method = "services?team_id="
	storeService         method = "services"
	updateService        method = "services/{id}"
	deleteService        method = "services/{id}"
	getLatestService     method = "services/latest"

	getEmployees        method = "employees?team_id="
	getMonthlyEmployees method = "monthly-employees?team_id="

	getTaxes method = "taxes"

	getTeams            method = "teams?team_id="
	getTeamIntegrations method = "teams/{id}/integrations"
	getTeamWidgetToggle method = "teams/{id}/widgets/toggle"
	getTeamRequsisites  method = "teams/{id}/requisites"
	getTeamImages       method = "teams/{id}/images"

	getTaxDepartments       method = "tax-departments"
	getSocialDepartments    method = "social-departments"
	getInsuranceDepartments method = "insurance-departments"

	storeBankStatement method = "bank-statements"

	getAccountRequisites method = "accounting-requisites"

	getBusinessTasks method = "business-tasks?team_id="

	getTickets           method = "tickets?team_id="
	storeTicket          method = "tickets"
	updateTicket         method = "tickets/{ticketId}"
	deleteTicket         method = "tickets/{ticketId}"
	getTicketComments    method = "tickets/{id}/comments?team_id="
	storeTicketComment   method = "tickets/{id}/comments"
	updateTicketComment  method = "tickets/{id}/comments/{commentId}"
	deleteTicketComment  method = "tickets/{id}/comments/{commentId}"
	getTicketDepartments method = "ticket-departments"

	getUnits    method = "units"
	getVatRates method = "vat-rates"

	getRoles     method = "roles"
	getConstants method = "constants"
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

func NewClient(baseUrl string, teamId string, token string) ClientApi {
	return ClientApi{
		c:         &http.Client{},
		endpoints: createEndpoints(baseUrl, teamId),
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

func createEndpoints(baseURI string, teamId string) map[method]string {
	list := make(map[method]string)

	list[getOperations] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getBankAccounts] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getBusinesses] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getContractors] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getCurrencies] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getDocuments] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getDocumentStatues] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getDocumentTypes] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getLatestOperation] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getLatestPayment] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getLatestService] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getOkvedCategories] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getOkveds] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getOperationFlows] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getOperationTypes] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getProfile] = fmt.Sprint(baseURI, "/", string(getProfile))
	list[getPayments] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getPaymentStatuses] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getServiceCategories] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getServices] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getLatestService] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getTaxes] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getTeams] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getTicketComments] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getTicketDepartments] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getTickets] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getUnits] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))
	list[getVatRates] = fmt.Sprint(baseURI, "/", string(getOperations), string(teamId))

	list[storeOperation] = fmt.Sprint(baseURI, "/", string(storeOperation))

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
