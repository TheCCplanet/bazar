package model

type Product struct {
	ID          int
	SellerID    int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Images      []ProductImage
}
type ProductImage struct {
	URL       string
	IsPrimary bool
}
