CREATE TABLE IF NOT EXISTS group_list (
	group_id INTEGER,
	user_id INTEGER,
	group_requested BOOLEAN,
	user_requested BOOLEAN,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (group_id) REFERENCES groups (id),
  	UNIQUE(user_id, group_id)
);