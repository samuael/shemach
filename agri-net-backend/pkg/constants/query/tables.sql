-- name: create-functionality-table
CREATE TABLE functionality_results (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    result boolean not null,
    reason varchar(300)
);
-- name: insert-functionality-results-table
INSERT INTO functionality_results (result, reason)
VALUES ($1, $2);

-- name: create-address-table
CREATE TABLE addresses (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    country varchar(50) NOT NULL,
    region VARCHAR(50) NOT NULL,
    zone VARCHAR(50) NOT NULL,
    woreda VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    kebele VARCHAR(50) NOT NULL
);
-- name: create-garage-table
CREATE TABLE garage (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name varchar(100) NOT NULL,
    address BIGINT REFERENCES addresses(id)
);
-- name: create-inspections-table
CREATE TABLE inspections (
    id BIGSERIAL primary key NOT NULL,
    garageid integer references garage(id),
    inspector_id integer REFERENCES functionality_results(id),
    drivername varchar(250) NOT NULL,
    vehicle_model varchar(100) NOT NULL,
    vehicle_year varchar(20) NOT NULL,
    vehicle_make varchar(100) NOT NULL,
    vehicle_color varchar(20) NOT NULL,
    license_plate varchar(100) UNIQUE NOT NULL,
    front_image varchar(100) NOT NULL,
    left_image varchar(100) NOT NULL,
    right_image varchar(100) NOT NULL,
    back_image varchar(100) NOT NULL,
    signature_image varchar(100) NOT NULL,
    vin_number varchar(100) UNIQUE NOT NULL,
    handbrake BIGINT NOT NULL,
    steering_system BIGINT NOT NULL,
    brake_system BIGINT NOT NULL,
    seat_belt BIGINT NOT NULL,
    door_and_window BIGINT NOT NULL,
    dashboard_light BIGINT NOT NULL,
    windshield BIGINT NOT NULL,
    baggage_door_window BIGINT NOT NULL,
    gear_box BIGINT NOT NULL,
    shock_absorber BIGINT NOT NULL,
    front_high_and_low_beam_light BIGINT NOT NULL,
    rear_light_and_brake_light BIGINT NOT NULL,
    wiper_operation BIGINT NOT NULL,
    car_horn BIGINT NOT NULL,
    side_mirror BIGINT NOT NULL,
    general_body_condition BIGINT NOT NULL,
    driver_performance boolean NOT NULL,
    balancing boolean NOT NULL,
    hazard boolean NOT NULL,
    signal_light_usage boolean NOT NULL,
    passed BOOLEAN NOT NULL
);
-- name: create-admins-table
CREATE TABLE admins (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    email varchar(200) unique not null,
    firstname varchar(100) not null,
    middlename varchar(100) not null,
    lastname varchar(100) not null,
    password Text not null,
    garageid integer references garage(id),
    inspectors_count integer default 0
);
-- name: create-inspectors-table
CREATE TABLE inspectors (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    email varchar(200) unique not null,
    firstname varchar(100) not null,
    middlename varchar(100) not null,
    lastname varchar(100),
    password text not null,
    imageurl varchar(30),
    garageid integer references garage(id),
    createdby integer references admins(id)
);
-- name: create-secretaries-table
CREATE TABLE secretaries (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    email varchar(200) unique not null,
    firstname varchar(100) not null,
    middlename varchar(100) not null,
    lastname varchar(100),
    password text not null,
    garageid integer references garage(id),
    createdby integer references admins(id)
);
-- name: insert-admin-table
INSERT INTO admins (
        email,
        firstname,
        middlename,
        lastname,
        password,
        garageid
    )
VALUES ($1, $2, $3, $4, $5, $6);
-- name: insert-garage-table
INSERT INTO garage (name, address)
VALUES ($1, $2);
-- name: insert-address-table
INSERT INTO addresses (
        country,
        region,
        zone,
        woreda,
        city,
        kebele
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6
    );