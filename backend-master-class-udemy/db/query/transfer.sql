-- name: CreateTransfer :one
INSERT INTO public.transfer (
    from_account_id,
    to_account_id,
    amount,
    created_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM public.transfer
WHERE transfer_id = $1 LIMIT 1; 

-- name: ListTransfers :many
SELECT * FROM public.transfer
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
UPDATE public.transfer
SET
    from_account_id = $1,
    to_account_id = $2,
    amount = $3,
    created_at = $4
WHERE transfer_id = $5
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM public.transfer
WHERE transfer_id = $1;