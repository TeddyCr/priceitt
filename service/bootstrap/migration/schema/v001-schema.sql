-- Create user table
CREATE TABLE IF NOT EXISTS "users" (
    json jsonb NOT NULL,
    id varchar(36) GENERATED ALWAYS AS (
        (json->>'id')::varchar
    ) STORED NOT null constraint users_id_pk PRIMARY KEY ,
    name varchar(256) GENERATED ALWAYS AS (
        (json->>'name')::varchar
    ) STORED NOT NULL,
    email varchar(256) GENERATED ALWAYS AS (
        (json->>'email')::varchar
    ) STORED NOT NULL constraint email_unique UNIQUE,
    "updatedAt" bigint GENERATED ALWAYS AS (
        (json->>'updatedAt')::bigint
    ) STORED NOT NULL,
    "createdAt" varchar(256) GENERATED ALWAYS AS (
        (json->>'createdAt')::bigint
    ) STORED NOT NULL
);


-- Create Token table
CREATE TABLE IF NOT EXISTS "tokens" (
    id varchar(36) GENERATED ALWAYS AS (
        (json->>'id')::varchar
    ) STORED NOT null constraint tokens_id_pk PRIMARY KEY ,
    json jsonb NOT NULL,
    name VARCHAR(36) GENERATED always as (json->>'name'::varchar ) stored null,
    userId varchar(36) GENERATED ALWAYS AS (
        (json->>'userId')::varchar
    ) STORED NOT NULL constraint tokens_user_id_fk REFERENCES "users"("id"),
    tokenType varchar(36) GENERATED ALWAYS AS (
        (json->>'tokenType')::varchar
    ) STORED NOT NULL,
    token varchar(512) GENERATED ALWAYS AS (
        (json->>'token')::varchar
    ) STORED NOT NULL,
    expirationDate varchar(512) GENERATED ALWAYS AS (
        (json->>'expirationDate')::bigint
    ) STORED NULL,
    ip varchar(36) GENERATED ALWAYS AS (
        (json->>'ip')::varchar
    ) STORED NULL,
    deviceId varchar(36) GENERATED ALWAYS AS (
        (json->>'deviceId')::varchar
    ) STORED NULL,
    "createdAt" bigint GENERATED ALWAYS AS (
        (json->>'createdAt')::bigint
    ) STORED NOT NULL
);

-- Create token blacklist table
CREATE TABLE IF NOT EXISTS "token_blacklist" (
    id varchar(36) GENERATED ALWAYS AS (
        (json->>'id')::varchar
    ) STORED NOT null constraint token_blacklist_id_pk PRIMARY KEY,
    json jsonb NOT NULL,
    userId varchar(36) GENERATED ALWAYS AS (
        (json->>'userId')::varchar
    ) STORED NOT NULL,
    token VARCHAR(512) GENERATED ALWAYS AS (
        (json->>'token')::varchar
    ) STORED NOT NULL,
    name VARCHAR(36) GENERATED always as (json->>'name'::varchar ) stored null
);

CREATE INDEX IF NOT EXISTS token_blacklist_token_type_idx ON token_blacklist (name);
CREATE INDEX IF NOT EXISTS token_blacklist_user_id_idx ON token_blacklist (userId);
