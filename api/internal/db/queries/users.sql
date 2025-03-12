-- name: CreateWebhookUser :exec
INSERT INTO users ( id, first_name, last_name, email, phone_number,profile_image_url, created_at, updated_at, external_id ) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: DeleteWebhookUser :exec
DELETE FROM users WHERE id = $1;