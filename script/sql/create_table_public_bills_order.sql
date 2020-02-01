CREATE TABLE bills_order
(
    id             uuid         NOT NULL DEFAULT uuid_generate_v4(),
    account_id     uuid         NOT NULL,
    type           text         NOT NULL,
    money          integer      NOT NULL,
    payment        text         NOT NULL,
    payment_number text         NOT NULL,
    created_at     timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at     timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at     timestamp(0),

    CONSTRAINT bills_order_pk PRIMARY KEY (id)
);
