create table if not exists tokens (
        id serial primary key,
        user_id uuid not null references users(id) on delete cascade,
        refresh_token text not null unique,
        role_id smallint,
        expires_at TIMESTAMPTZ not null,
        created_at TIMESTAMPTZ not null default NOW()
);

CREATE index idx_tokens_user_id on tokens(user_id);

