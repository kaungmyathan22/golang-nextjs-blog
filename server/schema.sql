CREATE TABLE
    IF NOT EXISTS users (
        id BIGSERIAL PRIMARY KEY,
        name text NOT NULL,
        password text NOT NULL,
        email text NOT NULL,
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS accounts (
        id BIGSERIAL PRIMARY KEY,
        lastLogInAt TIMESTAMP,
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS tokens (
        id BIGSERIAL PRIMARY KEY,
        token text NOT NULL,
        expiredAt TIMESTAMP NOT NULL,
        tokenType TEXT NOT NULL CHECK (
            tokenType IN (
                'refresh_token',
                'password_reset_token',
                'email_verification'
            )
        ),
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS categories (
        id BIGSERIAL PRIMARY KEY,
        title text,
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS posts (
        id BIGSERIAL PRIMARY KEY,
        title text,
        content text,
        createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        authorId BIGINT REFERENCES users (id),
        categoryId BIGINT REFERENCES categories (id)
    )