package models

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}
