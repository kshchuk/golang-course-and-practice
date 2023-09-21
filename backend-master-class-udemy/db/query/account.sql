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
SELECT * FROM  public.account
WHERE account_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAсcounts :many
SELECT * FROM public.account
WHERE owner = $1
ORDER BY account_id
LIMIT $2
OFFSET $3;

-- name: UpdateAccount :one
UPDATE public.account
SET
  owner = $2,
  balance = $3,
  currency = $4
WHERE account_id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE public.account
SET balance = balance + sqlc.arg(amount)
WHERE account_id = sqlc.arg(account_id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM public.account
WHERE account_id = $1;