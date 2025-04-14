-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_id UUID NOT NULL,
    user_id UUID NOT NULL,
    amount_paid DOUBLE PRECISION NOT NULL,
    payment_method TEXT NOT NULL,
    transaction_status TEXT NOT NULL DEFAULT 'PENDING',
    transaction_timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE receipts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID NOT NULL,
    user_id UUID NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    payment_method TEXT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_transaction FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS receipts;
DROP TABLE IF EXISTS transactions;
DROP EXTENSION IF EXISTS "uuid-ossp";
