CREATE TABLE
    IF NOT EXISTS payments (
        id uuid NOT NULL,
        reservation_id uuid NOT NULL,
        customer_id uuid NOT NULL,
        amount money NOT NULL,
        "method" varchar(50) NOT NULL,
        "date" timestamptz NOT NULL,
        status integer NOT NULL,
        CONSTRAINT pk_payments PRIMARY KEY (id),
        CONSTRAINT unq_payments_reservation_id UNIQUE (reservation_id),
        CONSTRAINT unq_payments_customer_id UNIQUE (customer_id)
    );

ALTER TABLE payments
ADD
    CONSTRAINT fk_payments_customers FOREIGN KEY (customer_id) REFERENCES customers(id);

ALTER TABLE payments
ADD
    CONSTRAINT fk_payments_reservation FOREIGN KEY (reservation_id) REFERENCES reservations(id);