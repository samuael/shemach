-- name: create-admin
CREATE TABLE subscriber (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(250) NOT NULL,
    Phone VARCHAR(250) NOT NULL UNIQUE,
    lang text NOT NULL,
    role smallint not null default 2,
    subscriptions smallint [] default array[]::smallint[];
);



create table tempo_subscriber(
    id serial primary key , 
    fullname VARCHAR(250) NOT NULL,
    Phone VARCHAR(250) NOT NULL UNIQUE,
    lang text NOT NULL,
    role smallint not null default 2,
    confirmation char(5) not null,
    unix integer not null
);


create table tempo_subscribers_login(
    id serial primary key,
    phone varchar(13) not null,
    confirmation char(5) not null,
    unix integer not null,
    trials smallint default 0
);

create table users(
    id serial primary key,
    firstname varchar(70)  not null,
    lastname varchar(70) not null,
    phone   varchar(13) unique  not null,
    email varchar(50)  unique not null,
    imageurl varchar(200) default "",
    created_at integer default ROUND(extract(epoch from now())),
    password text not null
);

create table infoadmin(
    messages_count integer default 0,
    created_by integer not null
) inherits(users);


create tempo_infoadmin (
    registration_second
);

create table superadmin(
    registered_admins integer default 0,
    registered_products integer default 0

) inherits(users);



-- Here we need the unit to be specified.
-- I have not decided on what and how the unit should be DECLARED

create table product(
    id serial primary key,
    name varchar(200) not null,
    production_area  varchar(200) not null,
    unit_id integer not null,
    current_price float default 0.0,
    created_by integer,
    created_at integer default ROUND(extract( epoch  from now())),
    last_updated_time integer default ROUND(extract( epoch  from now()))
);



create table messages (
    id serial primary key,
    targets integer[] not null default array[-1]::smallint[],
    lang varchar(5) not null,
    data text not null,
    created_by integer not null,
    created_at integer not null default ROUND( extract(epoch from now()))
);

create table admin(
    merchants_created integer default 0,
    
    stores integer default 0
) inherits(users);

