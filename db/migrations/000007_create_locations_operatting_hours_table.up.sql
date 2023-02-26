CREATE TABLE
    IF NOT EXISTS locations_operatting_hours (
        location_id uuid NOT NULL,
        day_of_week integer NOT NULL,
        opening_time time NOT NULL,
        closing_time time NOT NULL,
        CONSTRAINT pk_locations_operatting_hours PRIMARY KEY (location_id, day_of_week),
        CONSTRAINT unq_locations_operatting_hours_location_id UNIQUE (location_id)
    );

ALTER TABLE
    locations_operatting_hours
ADD
    CONSTRAINT fk_locations_operatting_hours FOREIGN KEY (location_id) REFERENCES locations(id);