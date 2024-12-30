-- Create user table
CREATE TABLE IF NOT EXISTS "users" (
    json jsonb NOT NULL,
    id varchar(36) GENERATED ALWAYS AS (
        (json->>'id')::varchar
    ) STORED NOT null constraint id_pk PRIMARY KEY ,
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
        (json->>'createdAt')::varchar
    ) STORED NOT NULL
);