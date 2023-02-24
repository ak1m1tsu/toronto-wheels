CREATE TABLE
    IF NOT EXISTS vehicles_images (
        id uuid NOT NULL,
        vehicle_id uuid NOT NULL,
        "path" varchar(400) NOT NULL,
        "primary" boolean NOT NULL,
        CONSTRAINT pk_vehicles_images PRIMARY KEY (id),
        CONSTRAINT unq_vehicles_images_vehicle_id UNIQUE (vehicle_id)
    );

ALTER TABLE vehicles
ADD
    CONSTRAINT fk_vehicles_vehicles_images FOREIGN KEY (id) REFERENCES vehicles_images(vehicle_id);