-- +goose Up
-- +goose StatementBegin
CREATE TABLE transaction_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    transaction_id UUID NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    status TEXT NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transaction_records;
-- +goose StatementEnd
