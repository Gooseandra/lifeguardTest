create table "user"(
    id serial primary key,
    name varchar(20) not null,
    surname varchar(30),
    patronymic varchar(20),
    password varchar(30) not null,
    nick varchar(20),
    phone varchar(10) not null,
    vk varchar(50),
    tg varchar(50),
    email varchar(50),
    private_data varchar(400)
);

create table "event"(
    id serial primary key,
    name varchar(45),
    discription varchar(100),
    order_discription varchar(400),
    start_date date,
    start_time time,
    end_date date,
    ned_time time
);

create table "event_roaster"(
    id serial primary key,
    userId integer references "user" (id),
    eventId integer references "event" (id),
    leader bool
);

create table "crew_inventory"(
    id serial primary key
);

create table "inventory"(
  id serial primary key,
  name varchar(50),
  value int,
  crew_inventory integer references "crew_inventory"(id)
);

create table "car"(
    id serial primary key,
    name varchar(40),
    comment varchar(400),
    inventory integer references crew_inventory(id)
);

create table "day_crew"(
    id serial primary key,
    time_start time,
    date_start date,
    time_end time,
    date_end time,
    comment varchar(400)
);

create table "date_crew_roaster"(
    "user" integer references "user"(id),
    day_crew integer references day_crew(id),
    leader bool
);

create table calls(
    id serial primary key,
    "desc" varchar(200),
    summing_up varchar(200),
    address varchar(45),
    time time
);

create table crew_calls(
    call_id integer references calls(id),
    crew_id integer references day_crew(id)
);

create table ranks(
  id serial primary key,
  name varchar(20)
);

create table permission(
    id serial primary key,
    action bool
);

create table "rank_permissions"(
    "permission" integer references "permission"("id"),
    "rank" integer references "ranks"("id")
);

create table "codes"(
    id serial primary key,
    code varchar(10),
    userId integer references "user"(id)
);

create table "event_inventory"(
    eventID integer references event(id),
    itemID integer references inventory(id)
);

create table "users_ranks"(
    rankId integer references ranks(id),
    userId integer references "user"(id)
);