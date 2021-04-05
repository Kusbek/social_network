CREATE TABLE IF NOT EXISTS groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    owner_id INTEGER,
    title VARCHAR,
    description VARCHAR,
    FOREIGN KEY (owner_id) REFERENCES users (id)
);