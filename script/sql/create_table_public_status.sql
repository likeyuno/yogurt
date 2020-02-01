CREATE TABLE status
(
    id               uuid         NOT NULL DEFAULT uuid_generate_v4(),
    node_id          uuid         NOT NULL,
    traffic_upload   interval     NOT NULL,
    traffic_download interval     NOT NULL,
    device           interval     NOT NULL,
    created_at       timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    updated_at       timestamp(0) NOT NULL DEFAULT current_timestamp(0),
    deleted_at       timestamp(0),

    CONSTRAINT status_pk PRIMARY KEY (id)
);
