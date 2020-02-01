CREATE TABLE packages
(
    id          uuid         NOT NULL DEFAULT uuid_generate_v4(),
    name        text         NOT NULL UNIQUE,
    description text         NOT NULL DEFAULT '',
    nodes       uuid[]       NOT NULL DEFAULT '{}',
    money       integer      NOT NULL,
    day         integer      NOT NULL,
    traffic     interval     NOT NULL,
    device      interval     NOT NULL,
    created_at  timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at  timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at  timestamp(0),

    CONSTRAINT packages_pk PRIMARY KEY (id)
);
