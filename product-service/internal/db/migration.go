package db

func RunMigrations() error {
	CreateProductTableQuerise := `
	CREATE TABLE IF NOT EXISTS products (
  id UUID PRIMARY KEY,
  seller_id UUID NOT NULL, -- user_id
  title VARCHAR(255),
  description TEXT,
  price NUMERIC(10, 2),
  quantity INT DEFAULT 1,
  image_url TEXT,
  category VARCHAR(100),
  is_active BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT NOW()
);`
	if _, err := DB.Exec(CreateProductTableQuerise); err != nil {
		return err
	}
	return nil
}
