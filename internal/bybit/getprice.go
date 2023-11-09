package bybit

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Bybit struct {
}

func (Bybit) GetPrice(payType, fiat, tradeType, transAmount, authMaker string) float64 {
	req := `
	{"userId":"",
	"tokenId":"USDT",
	"currencyId":"` + fiat + `",
	"payment":["` + payType + `"],
	"side":"` + tradeType + `",
	"size":"10",
	"page":"1",
	"amount":"` + transAmount + `",
	"authMaker":` + authMaker + `,
	"canTrade":true}`

	reqBody := strings.NewReader(req)

	resp, err := http.Post("https://api2.bybit.com/fiat/otc/item/online", "application/json", reqBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("incorrect status code %d", resp.StatusCode)
	}

	var response BybitResponse
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&response); err != nil {
		log.Fatal(err)
	}

	result, err := strconv.ParseFloat(response.Result.Items[0].Price, 64)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
