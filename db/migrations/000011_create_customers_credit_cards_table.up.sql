CREATE TABLE
    IF NOT EXISTS customer_credit_cards (
        customer_id uuid NOT NULL,
        card_id uuid NOT NULL,
        CONSTRAINT pk_customer_credit_cards PRIMARY KEY (customer_id, card_id),
        CONSTRAINT unq_customer_credit_cards_customer_id UNIQUE (customer_id),
        CONSTRAINT unq_customer_credit_cards_card_id UNIQUE (card_id)
    );

ALTER TABLE credit_cards
ADD
    CONSTRAINT fk_credit_cards FOREIGN KEY (id) REFERENCES customer_credit_cards(card_id);

ALTER TABLE customers
ADD
    CONSTRAINT fk_customers FOREIGN KEY (id) REFERENCES customer_credit_cards(customer_id);