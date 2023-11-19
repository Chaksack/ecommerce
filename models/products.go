package models

type Product struct {
	Id          uint     `json:"id"`
	Image       string   `json:"image"`
	ProductName string   `json:"product_name"`
	CategoryId  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:category_id"`
	Price       string   `json:"price"`
	Description string   `json:"description"`
	Quantity    int      `json:"quantity"`
}
