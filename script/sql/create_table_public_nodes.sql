CREATE TABLE nodes
(
    id          uuid         NOT NULL DEFAULT uuid_generate_v4(),
    name        text         NOT NULL UNIQUE,
    description text         NOT NULL DEFAULT '',
    tags        text[]       NOT NULL DEFAULT '{}',
    location    text         NOT NULL,
    type        text         NOT NULL,
    fast_open   bool         NOT NULL DEFAULT FALSE,
    created_at  timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at  timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at  timestamp(0),

    CONSTRAINT nodes_pk PRIMARY KEY (id)
);
