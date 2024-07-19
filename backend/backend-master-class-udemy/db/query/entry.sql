-- name: CreateEntry :one
INSERT INTO public.entry (
    account_id,
    amount,
    created_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM public.entry 
WHERE entry_id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM public.entry
LIMIT $1
OFFSET $2;

-- name: UpdateEntry :one
UPDATE public.entry
  set account_id = $2,
    amount = $3,
    created_at = $4
WHERE entry_id = $1
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM public.entry
WHERE entry_id = $1;