-- +goose Up
ALTER TABLE trading.asset ADD COLUMN description TEXT;

-- +goose Down
ALTER TABLE trading.asset DROP COLUMN description;
