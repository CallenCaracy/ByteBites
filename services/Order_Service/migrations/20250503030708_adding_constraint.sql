-- +goose Up
-- +goose StatementBegin
ALTER TABLE cart_items
ADD CONSTRAINT unique_cart_menu_item UNIQUE (cart_id, menu_item_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE cart_items
DROP CONSTRAINT IF EXISTS unique_cart_menu_item;
-- +goose StatementEnd
