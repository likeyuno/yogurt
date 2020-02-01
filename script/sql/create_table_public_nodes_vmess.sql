CREATE TABLE nodes_vmess
(
    id                uuid         NOT NULL DEFAULT uuid_generate_v4(),
    node_id           uuid         NOT NULL UNIQUE,
    host              text         NOT NULL,
    port              integer      NOT NULL,
    security          text         NOT NULL DEFAULT 'auto',
    alert_id          integer      NOT NULL,
    tls               bool         NOT NULL DEFAULT FALSE,
    tls_param         json         NOT NULL,
    obfuscation       text         NOT NULL,
    obfuscation_param json         NOT NULL,
    multiplex         bool         NOT NULL DEFAULT False,
    created_at        timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at        timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at        timestamp(0),

    CONSTRAINT nodes_vmess_pk PRIMARY KEY (id)
);
