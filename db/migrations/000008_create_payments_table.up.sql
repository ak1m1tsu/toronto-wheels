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

ALTER TABLE reservations
ADD
    CONSTRAINT fk_reservations_payments FOREIGN KEY (id) REFERENCES payments(reservation_id);