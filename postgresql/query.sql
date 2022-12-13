-- name: GetClientTypes :many
SELECT * FROM client_types ORDER BY id;

-- name: GetClients :many
SELECT * FROM clients ORDER BY name;

-- name: GetClientsByType :many
SELECT * FROM clients
WHERE client_type_id=$1 ORDER BY name;

-- name: CreateClient :one
INSERT INTO clients (
    name,
    client_type_id,
    tax_code, 
    address, 
    phone_number
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteClient :exec
DELETE FROM clients
WHERE id = $1;

-- name: UpdateClient :one
UPDATE clients
SET name=$2,
    client_type_id=$3,
    tax_code=$4,
    address=$5,
    phone_number=$6
WHERE id = $1
RETURNING *;