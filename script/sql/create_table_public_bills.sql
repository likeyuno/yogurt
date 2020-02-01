CREATE TABLE bills
(
    id             uuid         NOT NULL DEFAULT uuid_generate_v4(),
    type           text         NOT NULL,
    account_id     uuid         NOT NULL,
    gift_id        text         NOT NULL UNIQUE,
    money          integer      NOT NULL,
    payment        text         NOT NULL,
    payment_number text         NOT NULL,
    created_at     timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at     timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at     timestamp(0),

    CONSTRAINT bills_pk PRIMARY KEY (id)
);
