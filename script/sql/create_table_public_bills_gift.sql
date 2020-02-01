CREATE TABLE bills_gift
(
    id         uuid         NOT NULL DEFAULT uuid_generate_v4(),
    account_id uuid         NOT NULL DEFAULT '',
    type       text         NOT NULL,
    money      integer      NOT NULL,
    package_id uuid         NOT NULL,
    status     text         NOT NULL DEFAULT 'normal',
    created_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at timestamp(0),

    CONSTRAINT bills_gift_pk PRIMARY KEY (id)
);
