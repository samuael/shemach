CREATE TABLE addresses (
    id bigint PRIMARY KEY DEFAULT generate_unique_number(),
    location GEOMETRY(Point, 4326) not null,
    created_by bigint REFERENCES users(id),
    created_at bigint default extract(epoch from now())
);


CREATE INDEX idx_neighborhood_location ON neighborhoods(location);
CREATE INDEX idx_address_location ON addresses(location);
-- CREATE INDEX idx_address_neighborhood ON addresses(neighborhood_id);


create table tempo_registration_info(
    id bigint PRIMARY KEY DEFAULT generate_unique_number(),
    phone varchar(13) not null,
    code varchar(5) not null,
    created_at bigint default extract(epoch from now()),
    trials smallint default 0
);

create table verified_phones(
    id bigint PRIMARY KEY DEFAULT generate_unique_number(),
    phone varchar(13) not null,
    created_at bigint default extract(epoch from now())
);

create table users(
    id bigint PRIMARY KEY DEFAULT generate_unique_number(),
    firstname varchar(70)  not null,
    lastname varchar(70) not null,
    phone   varchar(13) unique  not null,
    email varchar(50) unique,
    telegram varchar(50) unique,
    profile_img_id bigint references img(id) ON DELETE SET NULL,
    created_at bigint default ROUND(extract(epoch from now())),
    password text not null,
    role smallint default 0,
    bio varchar(255)
);

create table forgot_password_shortcode (
    phone   varchar(13) unique  not null,
    last_sent bigint default extract(epoch from now()),
    created_at bigint default extract(epoch from now()),
    code varchar(6) not null,
    email_sent smallint default 0,
    phone_sent smallint default 0,
    telegram_sent smallint default 0,
    trials smallint default 0
);

create table session (
    id bigint PRIMARY KEY DEFAULT generate_unique_number(),
    userid bigint unique not null, 
    created_at bigint default extract(epoch from now()),
    token text not null
);

create table img (
    id bigint PRIMARY KEY DEFAULT generate_unique_number(),
    authorizations smallint default 0,
    resource varchar(200) not null,
    created_by bigint references users(id) not null,
    created_at bigint default extract(epoch from now())
);

create table blurred_img(
    id bigint PRIMARY KEY DEFAULT generate_unique_number(),
    source_id bigint references img(id) not null,
    resource varchar(200) not null
);

create table emailInConfirmation(
    id serial primary key,
    userid integer not null,
    new_email varchar(100) not null unique,
    is_new_account boolean default false,
    old_email varchar(100),
    created_at integer not null default ROUND(extract(epoch from now()))
);