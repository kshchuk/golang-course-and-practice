BEGIN;

ALTER TABLE IF EXISTS public.account DROP CONSTRAINT IF EXISTS owner_currency_key;

ALTER TABLE IF EXISTS public.account DROP CONSTRAINT IF EXISTS account_owner_fkey;

DROP TABLE IF EXISTS public.users;

END;