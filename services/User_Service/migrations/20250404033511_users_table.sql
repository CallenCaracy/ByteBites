-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD age INT DEFAULT NULL;

ALTER TABLE users
    ADD user_type TEXT DEFAULT NULL;

ALTER TABLE users
    ADD pfp TEXT DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN user_type;

ALTER TABLE users
    DROP COLUMN age;

ALTER TABLE users
    DROP COLUMN pfp;
-- +goose StatementEnd
