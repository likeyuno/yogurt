CREATE TABLE configs_rule
(
    id         uuid         NOT NULL DEFAULT uuid_generate_v4(),
    type       text         NOT NULL,
    value      text         NOT NULL,
    option     text         NOT NULL,
    remote_dns bool         NOT NULL DEFAULT FALSE,
    remote_tun bool         NOT NULL DEFAULT TRUE,
    script     bool         NOT NULL DEFAULT FALSE,
    created_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at timestamp(0),

    CONSTRAINT configs_rule_pk PRIMARY KEY (id)
);
