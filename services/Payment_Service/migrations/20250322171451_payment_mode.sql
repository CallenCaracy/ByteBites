-- +goose Up
-- +goose StatementBegin

CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
    amount DECIMAL(10, 2) NOT NULL CHECK (amount >= 0),
    payment_method VARCHAR(50) NOT NULL CHECK (payment_method IN ('Cash', 'Credit Card', 'Apple Pay', 'Google Pay', 'GCash')),
    status VARCHAR(20) NOT NULL DEFAULT 'Pending' CHECK (status IN ('Pending', 'Completed', 'Failed')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Business Logic: Trigger to validate payment amount matches order total
CREATE FUNCTION validate_payment_amount() RETURNS TRIGGER AS $$
BEGIN
    IF NEW.amount != (SELECT total_amount FROM orders WHERE order_id = NEW.order_id) THEN
        RAISE EXCEPTION 'Payment amount (%) must match order total (%)', NEW.amount, (SELECT total_amount FROM orders WHERE order_id = NEW.order_id);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_validate_payment_amount
    BEFORE INSERT OR UPDATE ON payments
    FOR EACH ROW
    EXECUTE FUNCTION validate_payment_amount();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS trigger_validate_payment_amount ON payments;
DROP FUNCTION IF EXISTS validate_payment_amount;
DROP TABLE payments;

-- +goose StatementEnd