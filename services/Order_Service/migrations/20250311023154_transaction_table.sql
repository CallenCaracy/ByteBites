-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    guest_session_id UUID REFERENCES guest_sessions(id) ON DELETE SET NULL,
    total_amount NUMERIC(10,2) NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('SUCCESSFUL','FAILED')),
    transaction_time TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
-- +goose StatementEnd
