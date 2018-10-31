package assistentby

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type method string

const (
	getProfile method = "user"

	getBankAccounts   method = "bank-accounts?team_id="
	storeBankAccount  method = "bank-accounts"
	updateBankAccount method = "bank-accounts/{id}"
	deleteBankAccount method = "bank-accounts/{id}"

	getBusinesses  method = "businesses?team_id="
	storeBusiness  method = "businesses"
	updateBusiness method = "businesses/{id}"
	deleteBusiness method = "businesses/{id}"

	getContractors   method = "contractors?team_id="
	storeContractor  method = "contractors"
	updateContractor method = "contractors/{id}"
	deleteContractor method = "contractors/{id}"

	getCurrencies method = "currencies"
	getBanks      method = "banks"

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

	updateTransactionStatus  method = "sync/transactions/status"
	updateBankAccountBalance method = "sync/bank-accounts/balance"
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

func NewClient(teamId string, token string, baseApiUrl string) ClientApi {
	if baseApiUrl == "" {
		baseApiUrl = BaseApiUrl
	}

	return ClientApi{
		c:         &http.Client{},
		endpoints: createEndpoints(baseApiUrl, teamId),
		baseUrl:   baseApiUrl,
		teamId:    teamId,
		token:     token,
	}
}

func (c *ClientApi) validateResponse(res *http.Response) ([]byte, error) {
	ct := res.Header.Get(http.CanonicalHeaderKey("Content-Type"))

	if ct != "application/json" {
		return nil, errors.New("response is not json")
	}

	if res.StatusCode > 201 {
		body, _ := ioutil.ReadAll(res.Body)

		log.Println(string(body))

		return nil, errors.New("validation response error")
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
	list[getBankAccounts] = fmt.Sprint(baseURI, "/", string(getBankAccounts), string(teamId))
	list[getBusinesses] = fmt.Sprint(baseURI, "/", string(getBusinesses), string(teamId))
	list[getContractors] = fmt.Sprint(baseURI, "/", string(getContractors), string(teamId))
	list[getCurrencies] = fmt.Sprint(baseURI, "/", string(getCurrencies))
	list[getBanks] = fmt.Sprint(baseURI, "/", string(getBanks))
	list[getDocuments] = fmt.Sprint(baseURI, "/", string(getDocuments), string(teamId))
	list[getDocumentStatues] = fmt.Sprint(baseURI, "/", string(getDocumentStatues), string(teamId))
	list[getDocumentTypes] = fmt.Sprint(baseURI, "/", string(getDocumentTypes), string(teamId))
	list[getLatestOperation] = fmt.Sprint(baseURI, "/", string(getLatestOperation), string(teamId))
	list[getLatestPayment] = fmt.Sprint(baseURI, "/", string(getLatestPayment), string(teamId))
	list[getLatestService] = fmt.Sprint(baseURI, "/", string(getLatestService), string(teamId))
	list[getOkvedCategories] = fmt.Sprint(baseURI, "/", string(getOkvedCategories), string(teamId))
	list[getOkveds] = fmt.Sprint(baseURI, "/", string(getOkveds), string(teamId))
	list[getOperationFlows] = fmt.Sprint(baseURI, "/", string(getOperationFlows), string(teamId))
	list[getOperationTypes] = fmt.Sprint(baseURI, "/", string(getOperationTypes), string(teamId))
	list[getProfile] = fmt.Sprint(baseURI, "/", string(getProfile))
	list[getPayments] = fmt.Sprint(baseURI, "/", string(getPayments), string(teamId))
	list[getPaymentStatuses] = fmt.Sprint(baseURI, "/", string(getPaymentStatuses), string(teamId))
	list[getServiceCategories] = fmt.Sprint(baseURI, "/", string(getServiceCategories), string(teamId))
	list[getServices] = fmt.Sprint(baseURI, "/", string(getServices), string(teamId))
	list[getLatestService] = fmt.Sprint(baseURI, "/", string(getLatestService), string(teamId))
	list[getTaxes] = fmt.Sprint(baseURI, "/", string(getTaxes), string(teamId))
	list[getTeams] = fmt.Sprint(baseURI, "/", string(getTeams), string(teamId))
	list[getTicketComments] = fmt.Sprint(baseURI, "/", string(getTicketComments), string(teamId))
	list[getTicketDepartments] = fmt.Sprint(baseURI, "/", string(getTicketDepartments), string(teamId))
	list[getTickets] = fmt.Sprint(baseURI, "/", string(getTickets), string(teamId))
	list[getUnits] = fmt.Sprint(baseURI, "/", string(getUnits), string(teamId))
	list[getVatRates] = fmt.Sprint(baseURI, "/", string(getVatRates), string(teamId))

	list[storeOperation] = fmt.Sprint(baseURI, "/", string(storeOperation))
	list[storePayment] = fmt.Sprint(baseURI, "/", string(storePayment))
	list[updateOperation] = fmt.Sprint(baseURI, "/", string(updateOperation))
	list[storeBankAccount] = fmt.Sprint(baseURI, "/", string(storeBankAccount))
	list[updateBankAccount] = fmt.Sprint(baseURI, "/", string(updateBankAccount))
	list[storeContractor] = fmt.Sprint(baseURI, "/", string(storeContractor))
	list[updateTransactionStatus] = fmt.Sprint(baseURI, "/", string(updateTransactionStatus))
	list[updateBankAccountBalance] = fmt.Sprint(baseURI, "/", string(updateBankAccountBalance))

	return list
}

func (c *ClientApi) postRequest(url string, postData interface{}) ([]byte, error) {
	j, _ := json.Marshal(postData)
	form := bytes.NewBuffer(j)

	req, err := http.NewRequest("POST", url, form)

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", c.token)

	response, err := c.c.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	res, err := c.validateResponse(response)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *ClientApi) putRequest(url string, putData interface{}) ([]byte, error) {
	j, _ := json.Marshal(putData)
	form := bytes.NewBuffer(j)
	log.Print(url, form)
	req, err := http.NewRequest("PUT", url, form)

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", c.token)

	response, err := c.c.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	res, err := c.validateResponse(response)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *ClientApi) getRequest(url string, params string) ([]byte, error) {
	req, err := http.NewRequest("GET", url+"?"+params, nil)

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", c.token)

	response, err := c.c.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	res, err := c.validateResponse(response)

	if err != nil {
		return nil, err
	}

	return res, err
}
