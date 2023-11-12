-- name: ListAssets :many
SELECT * from trading.asset;

-- name: GetAsset :one
SELECT * from trading.asset where asset_id = $1;

-- name: CreateAsset :one
INSERT INTO trading.asset (name, ticker, description) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateAsset :one
UPDATE trading.asset SET name = $1, ticker = $2, description = $3 WHERE asset_id = $4 RETURNING *;

-- name: DeleteAsset :exec
DELETE FROM trading.asset WHERE asset_id = $1;
