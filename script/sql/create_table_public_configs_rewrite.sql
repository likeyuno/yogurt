CREATE TABLE configs_rewrite
(
    id         uuid         NOT NULL DEFAULT uuid_generate_v4(),
    url        text         NOT NULL,
    value      text         NOT NULL,
    option     text         NOT NULL,
    script     bool         NOT NULL DEFAULT FALSE,
    created_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at timestamp(0),

    CONSTRAINT configs_rewrite_pk PRIMARY KEY (id)
);
