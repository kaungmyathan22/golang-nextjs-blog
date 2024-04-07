CREATE TABLE
  IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email text NOT NULL,
    fullName text,
    password text NOT NULL
  );

-- Add a unique constraint to the email column
ALTER TABLE users ADD CONSTRAINT email_unique UNIQUE (email);