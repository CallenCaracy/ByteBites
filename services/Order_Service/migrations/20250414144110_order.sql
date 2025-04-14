-- +goose Up
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL, 
    total_price DECIMAL(10,2) NOT NULL,
    order_status TEXT NOT NULL DEFAULT 'pending',
    order_type TEXT CHECK (order_type IN ('DINE_IN', 'TAKEAWAY', 'DELIVERY')),
    delivery_address TEXT DEFAULT NULL,
    special_requests TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE TABLE order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID REFERENCES orders(id) ON DELETE CASCADE,
    menu_item_id UUID NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    price DECIMAL(10,2) NOT NULL,
    customizations JSONB DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose Down
DROP TABLE orders;
DROP TABLE order_items;

