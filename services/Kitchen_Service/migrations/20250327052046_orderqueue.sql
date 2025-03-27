-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_queue (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL,
    status TEXT CHECK (status IN ('cooking', 'preparing', 'ready')) NOT NULL DEFAULT 'preparing',
    priority INT NOT NULL DEFAULT 1,
    last_updated TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_queue;
-- +goose StatementEnd
