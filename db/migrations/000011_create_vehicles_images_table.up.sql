CREATE TABLE
    IF NOT EXISTS vehicles_images (
        id uuid NOT NULL,
        vehicle_id uuid NOT NULL,
        "path" varchar(400) NOT NULL,
        "primary" boolean NOT NULL,
        CONSTRAINT pk_vehicles_images PRIMARY KEY (id),
        CONSTRAINT unq_vehicles_images_vehicle_id UNIQUE (vehicle_id)
    );

ALTER TABLE vehicles_images
ADD
    CONSTRAINT fk_vehicles_images_vehicles FOREIGN KEY (vehicle_id) REFERENCES vehicles(id);