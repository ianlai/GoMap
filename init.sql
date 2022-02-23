CREATE TABLE map
(
    id serial PRIMARY KEY,
    uid VARCHAR(20) NOT NULL UNIQUE,
    val INTEGER NOT NULL,
    updated_at TIMESTAMP
);