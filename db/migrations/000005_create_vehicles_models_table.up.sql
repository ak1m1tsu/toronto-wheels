CREATE TABLE
    IF NOT EXISTS vehicle_models (
        id uuid NOT NULL,
        name varchar(150) NOT NULL,
        make_id uuid NOT NULL,
        CONSTRAINT pk_vehicle_models PRIMARY KEY (id),
        CONSTRAINT unq_vehicle_models_make_id UNIQUE (make_id)
    );

ALTER TABLE vehicle_models
ADD
    CONSTRAINT fk_vehicle_models FOREIGN KEY (make_id) REFERENCES vehicle_makes(id);