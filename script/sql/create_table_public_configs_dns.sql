CREATE TABLE configs_dns
(
    id         uuid         NOT NULL DEFAULT uuid_generate_v4(),
    name       uuid         NOT NULL UNIQUE,
    host       text[]       NOT NULL,
    created_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at timestamp(0),

    CONSTRAINT configs_dns_pk PRIMARY KEY (id)
);
