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
    created_at integer default ROUND(extract(epoch from now())),
    password text not null
);

create table superadmin(
    registered_admins integer default 0,
    registered_products integer default 0
) inherits(users);



create table product(
    id serial primary key;
    name varchar(200) not null,
    production_area  varchar(200) not null,
    current_price float default 0.0,
    created_by integer,
    created_at integer default ROUND(extract( epoch  from now())),
    last_updated_time integer
);


create table admin(
    -- merchants_created integer default 0,
    -- stores
) inherits(user);



-- create table 
-- name: create-categories
-- CREATE TABLE category(
--     id serial primary key,
--     title varchar(50) not null unique,
--     short_title varchar(50) not null unique,
--     rounds_count integer default 0,
--     imgurl varchar(200),
--     fee decimal default 0.0,
--     created_at integer references eth_date(id)
-- );
-- -- name: create-round
-- CREATE TABLE round(
--     id serial primary key,
--     categoryid integer not null references category(id),
--     training_hour integer not null,
--     round_no integer not null UNIQUE check(round_no > 0),
--     students integer not null default 0,
--     active_amount decimal not null default 0.0,
--     active boolean default true,
--     start_date varchar(40),
--     lang char(3) default 'amh',
--     end_date varchar(40),
--     fee decimal default 0.0,
--     created_at integer references eth_date(id)
-- );
-- -- name: birth_date
-- CREATE TABLE eth_date (
--     id serial primary key,
--     year integer not null check(year >= 1920),
--     month integer not null check (
--         month >= 1
--         and month <= 13
--     ),
--     day integer not null check(
--         day >= 1
--         and day < 31
--     ),
--     hour integer default 0 check(
--         hour >= 0
--         and hour <= 24
--     ),
--     minute smallint default 0 check (
--         minute >= 0
--         and minute <= 60
--     ),
--     second smallint default 0 check (
--         second >= 0
--         and second <= 60
--     )
-- );
-- -- name: create-special-case
-- CREATE TABLE special_case (
--     id serial primary key,
--     reason text not null,
--     covered_amount decimal default 0.0,
--     complete_fee boolean default false
-- );
-- insert into special_case(id, reason)
-- values(0, '');
-- -- name: create-student
-- CREATE TABLE student (
--     id SERIAL PRIMARY KEY,
--     fullname VARCHAR(100) NOT NULL,
--     sex CHAR(1) NOT NULL DEFAULT 'M',
--     age decimal NOT NULL,
--     birth_date integer not null references eth_date(id),
--     accamic_status VARCHAR(100),
--     address integer references addresses(id),
--     phone varchar(15) unique not null,
--     paid decimal default 0.0,
--     status integer DEFAULT 1,
--     registered_by integer references admins(id),
--     round_id integer references round(id),
--     imgurl VARCHAR(200),
--     mark integer references special_case(id) default 0,
--     registered_at integer references eth_date(id)
-- );
-- -- name: create-addresses
-- CREATE TABLE addresses (
--     id SERIAL PRIMARY KEY,
--     city VARCHAR(60),
--     region VARCHAR(60),
--     zone varchar(40),
--     woreda varchar(40),
--     kebele VARCHAR(60),
--     unique_address varchar(60)
-- );
-- -- name: pay-in
-- CREATE TABLE payin (
--     id SERIAL PRIMARY KEY,
--     amount decimal not null,
--     recieved_by bigint references admins(id),
--     payed_by bigint references student(id),
--     created_at integer references eth_date(id),
--     roundid integer references round(id),
--     status smallint default 0,
--     uchars char(2) not null
-- );
-- -- name: create-payout 
-- CREATE TABLE payout(
--     id serial primary key,
--     title varchar(100) not null,
--     description text not null,
--     amount decimal not null default 0.0,
--     approved boolean default false,
--     withdrawed_by integer references admins(id),
--     created_at integer references eth_date(id),
--     status smallint default 0
-- );
-- -- name: insert-admin
-- INSERT into admins(
--         -- id, 
--         fullname,
--         email,
--         password,
--         superadmin,
--         imgurl
--     )
-- VALUES ($1, $2, $3, $4, $5);