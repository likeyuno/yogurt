CREATE TABLE nodes_shadosocks
(
    id                uuid         NOT NULL DEFAULT uuid_generate_v4(),
    node_id           uuid         NOT NULL UNIQUE,
    host              text         NOT NULL,
    port              integer      NOT NULL,
    password          text         NOT NULL,
    method            text         NOT NULL,
    obfuscation       text         NOT NULL,
    obfuscation_param json         NOT NULL,
    plugin            text         NOT NULL,
    plugin_param      json         NOT NULL,
    ota               bool         NOT NULL DEFAULT FALSE,
    created_at        timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at        timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at        timestamp(0),

    CONSTRAINT nodes_shadosocks_pk PRIMARY KEY (id)
);
