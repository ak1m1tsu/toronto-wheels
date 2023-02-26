CREATE TABLE
    IF NOT EXISTS vehicle_makes (
        id uuid NOT NULL,
        name varchar(150) NOT NULL,
        CONSTRAINT pk_vehicle_makes PRIMARY KEY (id)
    );