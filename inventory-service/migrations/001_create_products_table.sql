-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create index for better performance
CREATE INDEX IF NOT EXISTS idx_products_id ON products(id);
CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);


INSERT INTO products (name, stock) VALUES ('Product 1', 100);
INSERT INTO products (name, stock) VALUES ('Product 2', 200);
INSERT INTO products (name, stock) VALUES ('Product 3', 300);