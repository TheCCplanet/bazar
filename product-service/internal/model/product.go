package model

type Product struct {
	ID          int64
	SellerID    int64   `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	Images      []ProductImage
}
type ProductImage struct {
	URL       string
	IsPrimary bool
}
