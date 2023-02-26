CREATE TABLE
    IF NOT EXISTS customers (
        id uuid NOT NULL,
        first_name varchar(100) NOT NULL,
        last_name varchar(100) NOT NULL,
        email varchar(150) NOT NULL,
        phone varchar(50) NOT NULL,
        passhash varchar(1024) NOT NULL,
        license_number varchar(100) NOT NULL,
        address varchar(100) NOT NULL,
        CONSTRAINT pk_customers PRIMARY KEY (id)
    );

CREATE UNIQUE INDEX idx_customers_email ON customers (email);

CREATE UNIQUE INDEX idx_customers_phone ON customers (phone);

CREATE UNIQUE INDEX idx_customers_license_number ON customers (license_number);