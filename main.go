package main

import "fmt"

func main() {
	teamId := "2e14ba8bdc4d4f83916d0a8868b9d06f"
	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6Ijg0NmQzMjJlOWQ5NTgxZmUzMDRmMjY0ZWQwZTk1NWNjMmU5OGM3NDQxOWUwYzJjMzAwOTQwZWM4OTA0ZjAxNmRlZjdhODU2OTQ2MzVlZDIwIn0.eyJhdWQiOiIxIiwianRpIjoiODQ2ZDMyMmU5ZDk1ODFmZTMwNGYyNjRlZDBlOTU1Y2MyZTk4Yzc0NDE5ZTBjMmMzMDA5NDBlYzg5MDRmMDE2ZGVmN2E4NTY5NDYzNWVkMjAiLCJpYXQiOjE1Mzc2MjgzMzgsIm5iZiI6MTUzNzYyODMzOCwiZXhwIjoxODUzMjQ3NTM4LCJzdWIiOiIxOSIsInNjb3BlcyI6W119.W-YwbPlRz87N2NSNajyiPGveP98srtD6XQg7g180n2sXGHJM_QdL5UCXZzLyJWKjU0tcZMQUAdqJsUWYwHe9QQxuIrZaJ76qbQ2B8NWVrqkfygJW6dzJ27b5G4GOuSRLkZ8NeQP00k0k0OWMPWEBuPGiHrFJJ11I-6aUzT3yrDpZdYOFECKgeomTvMh5sdApCp4NIyrMa-jmb-NSFruUgfCqKLuwVC9yJTStMN4RwEt-thZ3L2XjT1tab5UYqcCR7rfn5CRIkXAvy0Ar1ItqQlE5f7-1_nGy7_li7WljXP9boQ7wmQKVX9Ac_n0FQfBPUB8aRJFBh3Ox-E3wvPmwmIcl1H-FXgm4Pt2R_287kFdjdaDIqoHad6kT40EL4uX47TnDhrgfUmZQEBQp7V9KDcI4kzpMVuoVYWM0T9N4cga1CPM0hVf4Go8CzAXtvCDdzf5KB1OKWxTT-6_A9anMY5Kcy1g2vAuapYi-I41QX9o5mxXVBBzaHpgwcAGsCaI14epEM3WOO3M1z1NpM2GX136o8YW_kOxHevXbi6rFrYIVHEy35KtU9u2fxC3aSWNtvnGssJLefAIWM1X6MCnVdYNKoFfR0AK7e61ExtaFvS5YmEkxe5nnMJ68GyWz71reRTSyJZOxTIAovMLaUjHgpuJg69Gx3FACCjZ7fWvZIwU"

	api := ApiAssistent{
		c: NewClient(teamId, token, ""),
	}

	params := "sort:id.asc"

	operations := api.getOperations(params)

	fmt.Printf("%+v\n", operations)
}
