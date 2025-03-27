-- +goose Up
-- +goose StatementBegin

CREATE TABLE cart_items (
    cart_item_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    item_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    customization TEXT,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL CHECK (subtotal >= 0),
    tax DECIMAL(10, 2) DEFAULT 0.00 CHECK (tax >= 0),
    delivery_fee DECIMAL(10, 2) DEFAULT 0.00 CHECK (delivery_fee >= 0),
    total_amount DECIMAL(10, 2) NOT NULL CHECK (total_amount >= 0),
    order_type VARCHAR(20) NOT NULL CHECK (order_type IN ('Delivery', 'Pickup')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
    item_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    price_at_time DECIMAL(10, 2) NOT NULL CHECK (price_at_time >= 0),
    customization TEXT
);

-- Business Logic: Trigger to calculate total_amount
CREATE FUNCTION calculate_order_total() RETURNS TRIGGER AS $$
BEGIN
    NEW.total_amount := NEW.subtotal + NEW.tax + NEW.delivery_fee;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_calculate_order_total
    BEFORE INSERT OR UPDATE ON orders
    FOR EACH ROW
    EXECUTE FUNCTION calculate_order_total();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS trigger_calculate_order_total ON orders;
DROP FUNCTION IF EXISTS calculate_order_total;
DROP TABLE order_items;
DROP TABLE orders;
DROP TABLE cart_items;

-- +goose StatementEnd