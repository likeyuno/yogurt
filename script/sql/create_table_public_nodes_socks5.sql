CREATE TABLE nodes_socks5
(
    id           uuid         NOT NULL DEFAULT uuid_generate_v4(),
    node_id      uuid         NOT NULL UNIQUE,
    host         text         NOT NULL,
    port         integer      NOT NULL,
    tls          bool         NOT NULL DEFAULT FALSE,
    tls_param    json         NOT NULL DEFAULT '',
    plugin       text         NOT NULL,
    plugin_param json         NOT NULL,
    created_at   timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at   timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at   timestamp(0),

    CONSTRAINT nodes_socks5_pk PRIMARY KEY (id)
);
