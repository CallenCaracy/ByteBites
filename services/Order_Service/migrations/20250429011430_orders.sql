-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL, 
    total_price DECIMAL(10,2) NOT NULL,
    order_status TEXT NOT NULL DEFAULT 'pending',
    order_type TEXT CHECK (order_type IN ('dine-in', 'takeout', 'delivery')),
    delivery_address TEXT DEFAULT NULL,
    special_requests TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
