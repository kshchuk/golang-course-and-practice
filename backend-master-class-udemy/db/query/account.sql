-- name: CreateAccount :one
INSERT INTO public.account (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM public.account 
WHERE account_id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAсcounts :many
SELECT * FROM public.account
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE public.account
SET
  owner = $2,
  balance = $3,
  currency = $4
WHERE account_id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM public.account
WHERE account_id = $1;