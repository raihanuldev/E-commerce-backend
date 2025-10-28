-- +migrate Up

CREATE TABLE IF NOT EXISTS
    users (
        id SERIAL PRIMARY KEY,
        frist_name VARCHAR(100) NOT NULL,
        last_name VARCHAR(100) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        is_shop_owner BOOLEAN NOT NULL DEFAULT FALSE,
        password TEXT NOT NULL,
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );