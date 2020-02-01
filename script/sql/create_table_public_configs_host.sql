CREATE TABLE configs_host
(
    id         uuid         NOT NULL DEFAULT uuid_generate_v4(),
    name       text UNIQUE  NOT NULL,
    host       text         NOT NULL,
    value      text         NOT NULL,
    created_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at timestamp(0),

    CONSTRAINT configs_host_pk PRIMARY KEY (id)
);
