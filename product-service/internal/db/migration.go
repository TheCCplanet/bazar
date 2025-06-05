package db

func RunMigrations() error {
	EnablePgcrypto := `CREATE EXTENSION IF NOT EXISTS "pgcrypto";`

	CreateProductTableQuery := `
	CREATE TABLE IF NOT EXISTS products (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		seller_id UUID NOT NULL, 
		title VARCHAR(255),
		description TEXT,
		price NUMERIC(10, 2),
		quantity INT DEFAULT 1,
		image_url TEXT,
		category VARCHAR(100),
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT NOW()
	);`

	CreateProductImageTableQuery := `
	CREATE TABLE IF NOT EXISTS product_images (
		id SERIAL PRIMARY KEY,
		product_id UUID REFERENCES products(id) ON DELETE CASCADE,
		image_url TEXT NOT NULL,
		is_primary BOOLEAN DEFAULT FALSE,
		position INTEGER DEFAULT 0
	);`

	CreateCategoryTableQuery := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description TEXT
	);`

	CreateAuctionTableQuery := `
	CREATE TABLE IF NOT EXISTS auctions (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		product_id UUID REFERENCES products(id) ON DELETE CASCADE,
		start_time TIMESTAMP NOT NULL,
		end_time TIMESTAMP NOT NULL,
		starting_price NUMERIC(15,2) NOT NULL,
		current_price NUMERIC(15,2),
		winner_id UUID
	);`

	CreateBidsTableQuery := `
	CREATE TABLE IF NOT EXISTS bids (
		id SERIAL PRIMARY KEY,
		auction_id UUID REFERENCES auctions(id) ON DELETE CASCADE,
		bidder_id UUID NOT NULL,
		bid_amount NUMERIC(15,2) NOT NULL,
		bid_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := DB.Exec(EnablePgcrypto); err != nil {
		return err
	}
	if _, err := DB.Exec(CreateProductTableQuery); err != nil {
		return err
	}
	if _, err := DB.Exec(CreateProductImageTableQuery); err != nil {
		return err
	}
	if _, err := DB.Exec(CreateCategoryTableQuery); err != nil {
		return err
	}
	if _, err := DB.Exec(CreateAuctionTableQuery); err != nil {
		return err
	}
	if _, err := DB.Exec(CreateBidsTableQuery); err != nil {
		return err
	}

	return nil
}
