CREATE TABLE subscriptions
(
    id         uuid         NOT NULL        DEFAULT uuid_generate_v4(),
    account_id uuid         NOT NULL,
    package_id uuid         NOT NULL,
    key        text         NOT NULL UNIQUE,
    uuid       text         NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    status     text         NOT NULL        DEFAULT 'normal',
    expired_at date         NOT NULL        DEFAULT current_date,
    created_at timestamp(0) NOT NULL        DEFAULT current_timestamp(0),
    updated_at timestamp(0) NOT NULL        DEFAULT current_timestamp(0),
    deleted_at timestamp(0),

    CONSTRAINT subscriptions_pk PRIMARY KEY (id)
);
