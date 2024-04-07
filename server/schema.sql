CREATE TABLE
  IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email text NOT NULL,
    fullName text,
    password text NOT NULL
  );