package binance

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Binance struct {
}

func (Binance) GetPrice(payType, fiat, tradeType, transAmount, merchant string) float64 {
	req := `
	{"proMerchantAds": false,
	"page": 1,
	"rows": 10,
	"payTypes": ["` + payType + `"],
	"countries": [],
	"publisherType": ` + merchant + `,
	"fiat": "` + fiat + `",
	"tradeType": "` + tradeType + `",
	"asset": "USDT",
	"merchantCheck": false,
	"transAmount": "` + transAmount + `"}`

	reqBody := strings.NewReader(req)

	resp, err := http.Post("https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search", "application/json", reqBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("incorrect status code %d", resp.StatusCode)
	}

	var response Response
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&response); err != nil {
		log.Fatal(err)
	}

	result, err := strconv.ParseFloat(response.Data[0].Adv.Price, 64)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
