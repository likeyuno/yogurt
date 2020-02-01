CREATE TABLE configs
(
    id         uuid         NOT NULL DEFAULT uuid_generate_v4(),
    account_id uuid         NOT NULL,
    key        text         NOT NULL UNIQUE,
    name       text         NOT NULL,
    dns        uuid         NOT NULL,

    created_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at timestamp(0),

    CONSTRAINT configs_pk PRIMARY KEY (id)
);
