CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR UNIQUE,
    email VARCHAR NOT NULL UNIQUE,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    birth_date TEXT NOT NULL,
    path_to_photo VARCHAR,
    about_me VARCHAR,
    is_public BOOLEAN DEFAULT 1,
    password VARCHAR NOT NULL
);