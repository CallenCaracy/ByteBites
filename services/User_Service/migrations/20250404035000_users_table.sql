-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD gender TEXT DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN gender;
-- +goose StatementEnd
