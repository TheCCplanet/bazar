package db

import (
	"bazar/internal/model"
	"fmt"
)

func AddProductWithImages(p *model.Product) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	err = tx.QueryRow(
		"INSERT INTO products (seller_id, title, description, price, quantitiy, category) VALUES($1,$2,$3,$4,$5,$6) RETURNING id, seller_id",
		p.SellerID, p.Title, p.Description, p.Price, p.Quantity, p.Category).Scan(p.ID)
	if err != nil {
		return err
	}
	for _, image := range p.Images {
		_, err = tx.Exec("INSERT INTO product_images (product_id, image_url, is_primary) VALUES($1,$2,$3)", p.ID, image.URL, image.IsPrimary)
	}
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Transaction commit failed %v", err)
	}
	// then add image url with productID database returned
	// with AddImageURL
	return nil
}
