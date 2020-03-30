CREATE DATABASE pseudonym;

GRANT ALL PRIVILEGES ON DATABASE pseudonym TO jameswu;

\connect pseudonym;

CREATE TABLE pseudonym
(
    id serial PRIMARY KEY,
    username VARCHAR,
    htmlurl VARCHAR,
    likes integer
);