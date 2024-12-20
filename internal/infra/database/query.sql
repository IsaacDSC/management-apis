-- name: GetAllEndpoints :many
SELECT * FROM endpoints WHERE deleted_at IS NULL and active = true order by created_at desc;

-- name: GetEndpoints :many
SELECT * FROM endpoints WHERE "service_name" = $1 and deleted_at IS NULL and active = true order by created_at desc;

-- name: GetServices :many
SELECT DISTINCT(service_name) FROM endpoints WHERE deleted_at IS NULL and active = true;

-- name: GetEndpoint :one
SELECT * FROM endpoints WHERE name = $1 and deleted_at IS NULL and active = true;

-- name: CreateOrUpdate :exec
INSERT INTO endpoints (service_name, name, description, method, url, path, headers, body, sensitive_api, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) ON CONFLICT (name) DO UPDATE SET service_name = $1, description = $3, method = $4, url = $5, path =  $6, headers = $7, body = $8, sensitive_api = $9, active = $10;

-- name: RemoveEndpoint :exec
UPDATE endpoints SET active = false, deleted_at = CURRENT_TIMESTAMP WHERE name = $1;

-- name: RemoveService :exec
UPDATE endpoints SET active = false, deleted_at = CURRENT_TIMESTAMP WHERE service_name = $1;


