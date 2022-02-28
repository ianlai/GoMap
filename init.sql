CREATE TABLE map
(
    id serial PRIMARY KEY,
    uid VARCHAR(20) NOT NULL UNIQUE,
    val INTEGER NOT NULL
);
CREATE INDEX index_val ON map (val);