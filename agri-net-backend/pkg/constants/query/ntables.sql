create table address (
    address_id serial primary key,
    kebele varchar(100),
    woreda varchar(100),
    city varchar(100),
    region varchar(100),
    unique_name  varchar(100),
    zone varchar(20),
    latitude varchar(20),
    longitude varchar(20)
);

-- name: create-admin
CREATE TABLE subscriber (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(250) NOT NULL,
    Phone VARCHAR(250) NOT NULL UNIQUE,
    lang text NOT NULL,
    role smallint not null default 2,
    subscriptions smallint [] default array[]::smallint[]
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
    imageurl varchar(200) default '',
    created_at integer default ROUND(extract(epoch from now())),
    password text not null,
    lang char(3) default 'amh'
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
    stores_created integer default 0,
    address_id integer references address(address_id),
    created_by integer not null
) inherits(users);


create table emailInConfirmation(
    id serial primary key,
    userid integer not null,
    new_email varchar(100) not null unique,
    is_new_account boolean default false,
    old_email varchar(100),
    created_at integer not null default ROUND(extract(epoch from now()))
);


create table agent (
    posts_count integer not null default 0,
    field_address_ref integer not null,
    registered_by  integer default 0 
) inherits( users );


create table merchant(
    stores integer default 0,
    posts_count  integer not null default 0,
    registerd_by integer not null,
    subscriptions smallint [] default array[]::smallint[]
    address_ref integer not null
) inherits ( users);

create table tempo_cxp (
    id serial primary key,
    phone varchar(13) not null unique,
    confirmation char(5) not null,
    role integer default 0,
    created_at integer not null default ROUND( extract(epoch from now())),
    trials integer default 0
);

-- alter table tempo_cxp add column trials integer default 0;


create table dictionary (
    id serial primary key,
    sentence text not null,
    lang char(3) not null, 
    translation text not null
)


create table store (
    store_id serial primary key,
    owner_id integer not null,
    address_id  int not null references address(address_id),
    active_products int default 0,
    store_name varchar(100),
    active_contracts int default 0,
    created_at  integer not null default ROUND(extract(  epoch from now())),
    created_by integer not null
);

alter table if exists store add constraint id_name_unique UNIQUE(store_id , store_name);


create table img (
    img_id serial primary key,
    resource varchar(200)   not null,
    owner_id integer not null,
    owner_role smallint not null,
    authorized boolean default false,
    authorizations smallint ,
    created_at integer not null default ROUND(extract(  epoch from now())),
    blurred_res   varchar(200)  not null
)


create table crop (
    crop_id serial primary key,
    type_id integer not null,
    description  varchar(500)  default '',
    negotiable boolean default false,
    remaining_quantity integer default 0,
    selling_price decimal not null default 0,
    address_id integer references address(address_id) not null,
    images  integer [] default array[]::integer[],
    created_at integer  not null default ROUND(extract(  epoch from now())),
    store_id integer references store(store_id),
    agent_id integer,
    store_owned boolean
);