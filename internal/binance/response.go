package binance

// type Response struct {
// 	Data []Data `json:"data"`
// }

// type Data struct {
// 	Adv Adv `json:"adv"`
// }

// type Adv struct {
// 	Price string `json:"price"`
// }

type Response struct {
	Data []struct {
		Adv struct {
			Price string `json:"Price"`
		}
	}
}
