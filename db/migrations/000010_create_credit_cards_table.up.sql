CREATE TABLE
    IF NOT EXISTS credit_cards (
        id uuid NOT NULL,
        "number" varchar(20) NOT NULL,
        expiration_date date NOT NULL,
        billing_address varchar(100) NOT NULL,
        CONSTRAINT pk_credit_cards PRIMARY KEY (id)
    );