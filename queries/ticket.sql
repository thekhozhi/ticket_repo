CREATE TABLE tickets(
ID uuid primary key not null,
from_city varchar(30) not null,
to_city varchar(30) not null,
date_of_flight TIMESTAMP not null
);

CREATE TABLE users(
    ID uuid primary key not null,
    first_name varchar(25) not null,
    last_name varchar(25) not null,
    email varchar(35),
    phone varchar(15) not null,
    ticket_id uuid references tickets(ID)
);
