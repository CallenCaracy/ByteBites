-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN age;

ALTER TABLE users
    ADD birthDate TIMESTAMPTZ
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    ADD age INT DEFAULT NULL;

ALTER TABLE users
    DROP COLUMN birth_date
-- +goose StatementEnd
