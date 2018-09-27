# assistentby-api

Example

    teamId := "qwerty123"
    token := "Bearer qwerty123"
    
	api := ApiAssistent{
		c: NewClient(teamId, token, ""),
	}

	params := "sort:id.asc"
	operations := api.getOperations(params)

	fmt.Printf("%+v\n", operations)