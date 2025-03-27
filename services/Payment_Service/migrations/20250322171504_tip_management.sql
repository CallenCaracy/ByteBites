-- +goose Up
-- +goose StatementBegin

CREATE TABLE tips (
    tip_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
    tip_amount DECIMAL(10, 2) NOT NULL CHECK (tip_amount >= 0),
    payment_method VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Business Logic: Trigger to update order total_amount when tip is added/updated
CREATE FUNCTION update_order_total_with_tip() RETURNS TRIGGER AS $$
BEGIN
    UPDATE orders
    SET total_amount = total_amount + NEW.tip_amount
    WHERE order_id = NEW.order_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_order_total_with_tip
    AFTER INSERT OR UPDATE ON tips
    FOR EACH ROW
    EXECUTE FUNCTION update_order_total_with_tip();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS trigger_update_order_total_with_tip ON tips;
DROP FUNCTION IF EXISTS update_order_total_with_tip;
DROP TABLE tips;

-- +goose StatementEnd