-- +goose Up
-- +goose StatementBegin
CREATE TABLE menu_list (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    category TEXT NOT NULL,
    availability_status BOOLEAN NOT NULL DEFAULT TRUE,
    image_url TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);
-- +goose StatementEnd 

-- +goose Down
-- +goose StatementBegin
DROP TABLE menu_list;
-- +goose StatementEnd
