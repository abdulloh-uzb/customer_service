CREATE TABLE if not exists admins (
  id bigserial PRIMARY KEY,
  email varchar NOT NULL,
  username varchar NOT NULL,
  password text NOT NULL,
  refresh_token text
)