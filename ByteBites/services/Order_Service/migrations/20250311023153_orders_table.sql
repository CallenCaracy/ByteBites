-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    guest_session_id UUID REFERENCES guest_sessions(id) ON DELETE CASCADE,
    menu_item_id UUID NOT NULL,  -- Menu items are stored in another DB
    quantity INT NOT NULL DEFAULT 1,
    order_time TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
