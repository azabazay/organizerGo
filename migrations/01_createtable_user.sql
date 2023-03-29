DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id serial PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    email VARCHAR(128) UNIQUE NOT NULL,
    password VARCHAR(128) NOT NULL
);

INSERT INTO
    users (name, email, password)
VALUES
    ('Abai Bekenov', 'azabazay@gmail.com', '12345');