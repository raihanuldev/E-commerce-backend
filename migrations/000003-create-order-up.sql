-- +migrate Up
CREATE TABLE
    IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        product_id BIGINT NOT NULL,
        user_id BIGINT NOT NULL,
        quantity INT NOT NULL,
        total_price DOUBLE PRECISION NOT NULL,
        status VARCHAR(20) NOT NULL DEFAULT 'pending',
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );