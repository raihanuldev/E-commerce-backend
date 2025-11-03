SELECT 
    id, product_id, user_id, quantity, total_price, status, created_at, updated_at
FROM orders
ORDER BY created_at DESC;
