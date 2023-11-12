-- +goose Up
CREATE SCHEMA trading;

CREATE TABLE trading.asset (
    asset_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    ticker TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE trading.price_history (
    price_history_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    asset_id uuid NOT NULL REFERENCES trading.asset(asset_id),
    price_asset_id uuid NOT NULL REFERENCES trading.asset(asset_id),
    price NUMERIC NOT NULL,
    event_time TIMESTAMPTZ NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE trading.price_history;
DROP TABLE trading.asset;
DROP SCHEMA trading;
