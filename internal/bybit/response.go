package bybit

type BybitResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	Items []Item `json:"items"`
}

type Item struct {
	Price string `json:"price"`
}
