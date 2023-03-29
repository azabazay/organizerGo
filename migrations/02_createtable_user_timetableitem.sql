DROP TABLE IF EXISTS time_table_items;

CREATE TABLE time_table_items (
  id serial PRIMARY KEY,
  user_id INT,
  time_start VARCHAR(128) NOT NULL,
  time_end VARCHAR(128) NOT NULL,
  name VARCHAR(128) NOT NULL,
  img_color VARCHAR(128) NOT NULL,
  img_url VARCHAR(128) NOT NULL,
  like_count INT
);

INSERT INTO
  time_table_items (
    user_id,
    time_start,
    time_end,
    name,
    img_color,
    img_url,
    like_count
  )
VALUES
  (
    1,
    '4:00pm',
    '5:00pm',
    'Contemprorary Dance',
    'FFB6C1',
    '',
    11
  ),
  (
    1,
    '5:00pm',
    '6:00pm',
    'Break Dance',
    '00FFFF',
    '',
    28
  ),
  (
    1,
    '5:00pm',
    '6:00pm',
    'Street Dance',
    '8A2BE2',
    '',
    28
  ),
  (1, '7:00pm', '8:00pm', 'Yoga', '6495ED', '', 23),
  (
    1,
    '6:00pm',
    '7:00pm',
    'Stretching',
    '00FFFF',
    '',
    14
  ),
  (
    1,
    '8:00pm',
    '9:00pm',
    'Street Dance',
    '008B8B',
    '',
    9
  );