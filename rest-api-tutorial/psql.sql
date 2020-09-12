CREATE TABLE person(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    gender varchar(7) NOT NULL,
    data_of_birth DATE NOT NULL,
    email varchar(150),
    country_of_birth varchar(50)
)

