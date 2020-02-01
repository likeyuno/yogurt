CREATE TABLE nodes_shadowsocksr
(
    id                uuid         NOT NULL DEFAULT uuid_generate_v4(),
    node_id           uuid         NOT NULL UNIQUE,
    host              text         NOT NULL,
    port              integer      NOT NULL,
    password          text         NOT NULL,
    method            text         NOT NULL,
    protocol          text         NOT NULL,
    protocol_param    text         NOT NULL,
    obfuscation       text         NOT NULL,
    obfuscation_param json         NOT NULL,
    plugin            text         NOT NULL,
    plugin_param      json         NOT NULL,
    created_at        timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at        timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at        timestamp(0),

    CONSTRAINT nodes_shadowsocksr_pk PRIMARY KEY (id)
);
