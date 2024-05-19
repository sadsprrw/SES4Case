package models

type CurrencyResponce struct {
	Rates CurrencyRates `json:"rates"`
}

type CurrencyRates struct {
	Currency float64 `json:"UAH"`
}
