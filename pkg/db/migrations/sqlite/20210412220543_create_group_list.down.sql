CREATE TABLE IF NOT EXISTS followers (
	user_id INTEGER,
  	group_id INTEGER,
	group_requested BOOLEAN DEFAULT 0,
	user_requested BOOLEAN DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (group_id) REFERENCES groups (id),
  	UNIQUE(user_id, following_id)
);