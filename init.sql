CREATE TABLE map
(
    uid serial PRIMARY KEY,
    id VARCHAR(20) NOT NULL,
    val INTEGER NOT NULL,
    updated_at TIMESTAMP
);