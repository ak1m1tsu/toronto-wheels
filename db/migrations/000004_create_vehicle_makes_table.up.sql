CREATE TABLE
    IF NOT EXISTS vehicle_makes (
        id uuid NOT NULL,
        name varchar(150) NOT NULL,
        CONSTRAINT pk_vehicle_makes PRIMARY KEY (id)
    );

ALTER TABLE vehicle_makes
ADD
    CONSTRAINT fk_vehicle_makes FOREIGN KEY (id) REFERENCES vehicle_models(make_id);

ALTER TABLE vehicle_makes
ADD
    CONSTRAINT fk_vehicle_makes_vehicles FOREIGN KEY (id) REFERENCES vehicles(make_id);