CREATE TABLE schools(
    id serial PRIMARY KEY,
    name varchar(255) UNIQUE
);

CREATE TABLE religion_details(
    id serial PRIMARY KEY,
    religion varchar(255) UNIQUE
);