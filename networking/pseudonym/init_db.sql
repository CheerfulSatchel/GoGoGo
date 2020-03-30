CREATE DATABASE pseudonym;

GRANT ALL PRIVILEGES ON DATABASE pseudonym_yeet TO jameswu;

\connect pseudonym;

CREATE TABLE pseudonym
(
    id integer PRIMARY KEY,
    username VARCHAR,
    html_url VARCHAR,
    likes integer
);