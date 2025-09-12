-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    version INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create index for better performance
CREATE INDEX IF NOT EXISTS idx_products_id ON products(id);
CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);

-- Insert dummy data
INSERT INTO products (name, stock) VALUES 
('Product 1', 100),
('Product 2', 200),
('Product 3', 300),
('Laptop Gaming', 50),
('Smartphone', 150),
('Headphones', 200),
('Keyboard Mechanical', 75),
('Mouse Wireless', 120),
('Monitor 4K', 30),
('Webcam HD', 80),
('Tablet', 90),
('Smart Watch', 60),
('Power Bank', 100),
('USB Cable', 500),
('Charger Fast', 200),
('Bluetooth Speaker', 85),
('Gaming Chair', 25),
('Desk Lamp', 150),
('External SSD', 40),
('Memory Card', 300),
('Phone Case', 250),
('Screen Protector', 400),
('Laptop Stand', 35);