-- name: CreateUser :one
INSERT INTO public.users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM public.users 
WHERE username = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM  public.users
WHERE username = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM public.users
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE public.users
SET
    hashed_password = $2,
    full_name = $3,
    email = $4,
    password_changed_at = $5
WHERE username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM public.users
WHERE username = $1;