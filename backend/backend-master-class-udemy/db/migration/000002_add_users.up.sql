BEGIN;

CREATE TABLE
    IF NOT EXISTS public.users (
        username VARCHAR(255) NOT NULL,
        hashed_password VARCHAR(255) NOT NULL,
        full_name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        created_at timestamp with time zone NOT NULL DEFAULT NOW(),
        password_changed_at timestamp with time zone DEFAULT NULL,
        CONSTRAINT username_pkey PRIMARY KEY (username)
    );

ALTER TABLE public.account
ADD FOREIGN KEY (owner) REFERENCES public.users (username) NOT VALID;

ALTER TABLE public.account
ADD CONSTRAINT owner_currency_key UNIQUE (owner, currency);

END;