CREATE TABLE
    IF NOT EXISTS locations (
        id uuid NOT NULL,
        address varchar(150) NOT NULL,
        CONSTRAINT pk_locations PRIMARY KEY (id)
    );