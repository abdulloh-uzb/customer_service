CREATE TABLE if not exists customers (
  id bigserial PRIMARY KEY,
  first_name varchar NOT NULL,
  last_name varchar NOT NULL,
  bio text NOT NULL,
  email varchar NOT NULL,
  password text NOT NULL,
  phone_number varchar NOT NULL,
  created_at timestamptz NULL DEFAULT now(),
  deleted_at timestamptz NULL,
  updated_at timestamptz NULL,
  refresh_token text
)