-- +goose Up
-- +goose StatementBegin
-- Create tables for the Payment Service

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE food_items (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    total_amount DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INTEGER REFERENCES carts(id) ON DELETE CASCADE,
    food_item_id INTEGER REFERENCES food_items(id) ON DELETE RESTRICT,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    cart_id INTEGER UNIQUE REFERENCES carts(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    status TEXT DEFAULT 'pending' CHECK (status IN ('pending', 'completed', 'cancelled')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    order_id INTEGER UNIQUE REFERENCES orders(id) ON DELETE CASCADE,
    payment_method TEXT NOT NULL CHECK (payment_method IN ('cash', 'credit_card', 'apple_pay', 'google_pay')),
    payment_info TEXT,
    amount_paid DECIMAL(10, 2) NOT NULL,
    transaction_id TEXT UNIQUE NOT NULL,
    transaction_status TEXT NOT NULL DEFAULT 'pending' CHECK (transaction_status IN ('pending', 'completed', 'failed')),
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE receipts (
    id SERIAL PRIMARY KEY,
    transaction_id TEXT UNIQUE NOT NULL REFERENCES payments(transaction_id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    order_id INTEGER REFERENCES orders(id) ON DELETE SET NULL,
    amount_paid DECIMAL(10, 2) NOT NULL,
    payment_method TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    transaction_status TEXT NOT NULL
);

-- Indexes for performance
CREATE INDEX idx_carts_user_id ON carts(user_id);
CREATE INDEX idx_cart_items_cart_id ON cart_items(cart_id);
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_payments_order_id ON payments(order_id);
CREATE INDEX idx_receipts_user_id ON receipts(user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop tables in reverse order to respect foreign key constraints

DROP TABLE receipts;
DROP TABLE payments;
DROP TABLE orders;
DROP TABLE cart_items;
DROP TABLE carts;
DROP TABLE food_items;
DROP TABLE users;

-- +goose StatementEnd