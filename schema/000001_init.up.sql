create table restaurants
(
    id serial not null unique,
    name varchar(100) not null unique,
    tables_count int not null,
    average_wait_time_minutes int,
    average_check_rubles int,
    work_day_start_time time not null default '09:00',
    work_day_end_time time not null default '23:00',
    booking_time_minutes int not null default 120,
    check ( work_day_end_time > work_day_start_time )
);

create table tables
(
    id serial not null unique,
    seats_count int not null,
    restaurant_id int references restaurants (id) on delete cascade not null
);

create table bookings
(
    id serial not null unique,
    user_name varchar(100) not null,
    user_phone varchar(50) not null,
    start_time time not null,
    booking_date date not null default CURRENT_DATE
);


-- this table is needed because there can be several tables in one booking
create table tables_bookings
(
    id serial not null unique,
    table_id int references tables (id) on delete cascade not null,
    booking_id int references bookings (id) on delete cascade not null
);
