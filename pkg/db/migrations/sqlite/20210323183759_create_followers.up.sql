CREATE TABLE IF NOT EXISTS followers (
	user_id INTEGER,
  	following_id INTEGER,
	is_requested BOOLEAN DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (following_id) REFERENCES users (id),
  	UNIQUE(user_id, following_id)
);