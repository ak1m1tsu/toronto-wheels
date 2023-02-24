CREATE TABLE
    IF NOT EXISTS locations (
        id uuid NOT NULL,
        address varchar(150) NOT NULL,
        CONSTRAINT pk_locations PRIMARY KEY (id)
    );

ALTER TABLE locations
ADD
    CONSTRAINT fk_locations_vehicles FOREIGN KEY (id) REFERENCES vehicles(location_id);

ALTER TABLE locations
ADD
    CONSTRAINT fk_locations_reservations_pickup FOREIGN KEY (id) REFERENCES reservations(pickup_location_id);

ALTER TABLE locations
ADD
    CONSTRAINT fk_locations_reservations_dropoff FOREIGN KEY (id) REFERENCES reservations(dropoff_location_id);