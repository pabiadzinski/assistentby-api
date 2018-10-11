# assistentby-api

package main

import (
	"fmt"
	a "github.com/pabiadzinski/assistentby-api"
)

func main() {
	teamId := "2e14ba8bdc4d4f83916d0a8868b9d06f"
	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjEyNDVmODNiZjdjMzMwNTBlYTA5MjY0MThmZjJiYjQ1ZDhhNDk3ZWY4YzZkYWY2MmYzYTE1ZDE1YzY0MjA0NGU2MGM1M2M4OTk4YWE3MWQ5In0.eyJhdWQiOiIxIiwianRpIjoiMTI0NWY4M2JmN2MzMzA1MGVhMDkyNjQxOGZmMmJiNDVkOGE0OTdlZjhjNmRhZjYyZjNhMTVkMTVjNjQyMDQ0ZTYwYzUzYzg5OThhYTcxZDkiLCJpYXQiOjE1Mzg0Njk1MjUsIm5iZiI6MTUzODQ2OTUyNSwiZXhwIjoxODU0MDg4NzI0LCJzdWIiOiIxOSIsInNjb3BlcyI6W119.JcMn4h6l9_j0mniBpDxWjnLppjVVyva_BCIwU1CaBsGmj3WR5HDEr1lHcydr9BJDiLElK3zmHw_Mjxi15ndkqii4vlU1qjKdI-JfwED6bMivsMnUBkRpKm1eOHxsfuYKBRG4nl4bAbGkGyqAEySPiOPzyfS-UFgtFh4KfLSF2y4l6kX59k4kRIjW2vyCXIvggAXPCoxttDPs6OvXaMToOiYAR-n2x-lhlZvqJrAJ3AxUba27IhdyOq3hJhVeLsqSBGlr2EIvrgAxS0kMdnzxTJ1_Mobndm6S_Tktk5U5yGVq-DWq0vXRiELcCF9hoNxdl97_a61JLt9_Ku_lt8Bw1r14D3lu-HVfn0kh3LjOcBEGcsuDcDvEdDBs9U6wwGQLbzhHCYbNE7ibvJ7XMZ4IjMmLIcBBG7FomGBBvO_zwnEw6pRY8UwQT8KzymM5p3xPjOELcAMVcbiyqMx7FuFm7qMjSvjl2fY7jYfJM4msnuJT002tXhWpwIFJaBI6rIDnf3rPWa0jzc9LBxqLmPNS1mkC5SXtn4MX-r1LJdHZRfdQU2K6zJKFjsaVgZtphHAP6qbq_RLE-rZGxnH91ylEYzXSJCwZ3YY6YGdiFlOeJT33tEnP4K7cnJlXH9_Idde6Cg64T9iATWr_ol99cMSJn84k0gFOc5OBc6kWu5vUFu4"

	api := a.ApiAssistent{
		C: a.NewClient(teamId, token, "http://app.local/api/v1"),
	}

	formBankAccount := a.BankAccount{
		TeamId:             teamId,
		BankId:             3,
		CurrencyId:         1,
		Number:             "BY97ALFA3012215177006237",
		OpeningBalance:     10,
		OpeningBalanceDate: "2018-09-01 00:00:00",
	}

	formOperation := a.Operation{
		TeamId:          teamId,
		Date:            "2018-09-01 00:00:00",
		AmountIn:        100,
		PaymentDocument: "777",
		Description:     "Описание платежа",
		OperationTypeId: a.OperationIncomeID,
		OperationFlowId: a.OperationFlowBankID,
		CurrencyId:      8,
		TransactionId:   "123",
	}

	responseBankAccounts := api.StoreBankAccount(formBankAccount)
	fmt.Printf("%+v\n", responseBankAccounts)

	respOperations := api.StoreOperation(formOperation)
	fmt.Printf("%+v\n", respOperations)
}
