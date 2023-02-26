CREATE TABLE
    IF NOT EXISTS customers_credit_cards (
        customer_id uuid NOT NULL,
        card_id uuid NOT NULL,
        CONSTRAINT pk_customer_credit_cards PRIMARY KEY (customer_id, card_id),
        CONSTRAINT unq_customer_credit_cards_customer_id UNIQUE (customer_id),
        CONSTRAINT unq_customer_credit_cards_card_id UNIQUE (card_id)
    );

ALTER TABLE
    customers_credit_cards
ADD
    CONSTRAINT fk_customer_credit_cards_card FOREIGN KEY (card_id) REFERENCES credit_cards(id);

ALTER TABLE
    customers_credit_cards
ADD
    CONSTRAINT fk_customer_credit_cards_customer FOREIGN KEY (customer_id) REFERENCES customers(id);