CREATE TABLE nodes_http
(
    id           uuid         NOT NULL DEFAULT uuid_generate_v4(),
    node_id      uuid         NOT NULL UNIQUE,
    host         text         NOT NULL,
    port         integer      NOT NULL,
    tls          bool         NOT NULL DEFAULT FALSE,
    tls_param    json         NOT NULL DEFAULT '',
    plugin       text         NOT NULL,
    plugin_param json         NOT NULL,
    http2        bool         NOT NULL DEFAULT FALSE,
    created_at   timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at   timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at   timestamp(0),

    CONSTRAINT nodes_http_pk PRIMARY KEY (id)
);
