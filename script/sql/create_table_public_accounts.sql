CREATE TABLE accounts
(
    id         uuid         NOT NULL        DEFAULT uuid_generate_v4(),
    username   text         NOT NULL UNIQUE,
    nickname   text         NOT NULL UNIQUE DEFAULT '',
    email      text         NOT NULL UNIQUE,
    qq         integer      NOT NULL UNIQUE DEFAULT 0,
    telegram   integer      NOT NULL UNIQUE DEFAULT 0,
    money      integer      NOT NULL        DEFAULT 0,
    created_at timestamp(0) NOT NULL        DEFAULT current_timestamp(0),
    updated_at timestamp(0) NOT NULL        DEFAULT current_timestamp(0),
    deleted_at timestamp(0),

    CONSTRAINT accounts_pk PRIMARY KEY (id)
);
