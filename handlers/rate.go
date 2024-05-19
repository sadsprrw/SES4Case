package handlers

import (
	"bytes"
	"encoding/json"
	"example/SES4Case/models"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FetchRate() (float64, error) {
	api_key := os.Getenv("API_KEY")
	url := "https://api.currencybeacon.com/v1/latest?api_key=" + api_key + "&base=USD&symbols=UAH"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	r := bytes.NewReader(body)
	decoder := json.NewDecoder(r)

	val := &models.CurrencyResponce{}
	err = decoder.Decode(val)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	rate := val.Rates.Currency
	return rate, nil
}

func GetRate(c *gin.Context) {
	rate, err := FetchRate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate fetch error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rate": rate})
}
