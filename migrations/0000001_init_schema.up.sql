CREATE TABLE IF NOT EXISTS urls (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    alias VARCHAR(255) NOT NULL UNIQUE,
    url TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_alias on urls(alias);