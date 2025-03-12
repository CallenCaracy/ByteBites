-- +goose Up
-- +goose StatementBegin
CREATE TABLE guest_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_token TEXT UNIQUE NOT NULL,
    started_at TIMESTAMP DEFAULT NOW(),
    ended_at TIMESTAMP,                  
    status TEXT NOT NULL DEFAULT 'active'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE guest_sessions;
-- +goose StatementEnd
