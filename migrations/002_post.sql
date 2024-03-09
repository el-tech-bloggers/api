-- Write your migrate up statements here

CREATE TABLE post(
  id serial primary key,
  title VARCHAR(150) NOT NULL,
  body TEXT NOT NULL
);


INSERT INTO post(title, body) VALUES
('title1', 'description1'),
('title2', 'description2'),
('title3', 'description3');

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible

DROP TABLE post;

-- Then delete the separator line above.
