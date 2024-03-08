-- Write your migrate up statements here

CREATE TABLE admin(
  id serial primary key,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(50) NOT NULL
);

INSERT INTO admin(username, password) VALUES ('abdallah', '{{.admin1Password}}');
INSERT INTO admin(username, password) VALUES ('bassel', '{{.admin2Password}}');

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible

DROP TABLE admin;

-- Then delete the separator line above.
