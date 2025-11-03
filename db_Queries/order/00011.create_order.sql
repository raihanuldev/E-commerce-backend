INSERT INTO orders (product_id, user_id, quantity, total_price, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
RETURNING id;
