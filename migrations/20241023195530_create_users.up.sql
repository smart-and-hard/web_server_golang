CREATE TABLE users (
    id bigserial not null PRIMARY KEY,
    email varchar not null UNIQUE,
    encrypted_password varchar not null 
);