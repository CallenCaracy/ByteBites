-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventory (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_name TEXT NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0),
    unit TEXT NOT NULL,
    low_stock_threshold INT NOT NULL DEFAULT 5,
    expiry_date TIMESTAMPTZ DEFAULT NULL,
    last_updated TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inventory;
-- +goose StatementEnd
