CREATE TABLE
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

SELECT
    id,
    frist_name,
    last_name,
    email,
    is_shop_owner,
    password
FROM
    users
ORDER BY
    id

INSERT INTO users (frist_name, last_name, email, is_shop_owner, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
SELECT id, frist_name, last_name, email, is_shop_owner, password
		FROM users
		WHERE email = $1