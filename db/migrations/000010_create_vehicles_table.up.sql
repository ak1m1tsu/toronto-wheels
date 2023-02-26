CREATE TABLE
    IF NOT EXISTS vehicles (
        id uuid NOT NULL,
        make_id uuid NOT NULL,
        model_id uuid NOT NULL,
        "year" integer NOT NULL,
        location_id uuid NOT NULL,
        "type" integer NOT NULL,
        availability boolean NOT NULL,
        fuel_type integer NOT NULL,
        transmission_type integer NOT NULL,
        price_per_day money NOT NULL,
        CONSTRAINT pk_vehicles PRIMARY KEY (id),
        CONSTRAINT unq_vehicles_model_id UNIQUE (model_id),
        CONSTRAINT unq_vehicles_make_id UNIQUE (make_id),
        CONSTRAINT unq_vehicles_location_id UNIQUE (location_id)
    );

CREATE INDEX idx_vehicles_type ON vehicles ( "type" );

CREATE INDEX
    idx_vehicles_transmission_type ON vehicles (transmission_type);

CREATE INDEX idx_vehicles_fuel_type ON vehicles (fuel_type);

ALTER TABLE vehicles
ADD
    CONSTRAINT fk_vehicles_vehicle_makes FOREIGN KEY (make_id) REFERENCES vehicle_makes(id);

ALTER TABLE vehicles
ADD
    CONSTRAINT fk_vehicles_vehicle_models FOREIGN KEY (model_id) REFERENCES vehicle_models(id);

ALTER TABLE vehicles
ADD
    CONSTRAINT fk_vehicles_locations FOREIGN KEY (location_id) REFERENCES locations(id);