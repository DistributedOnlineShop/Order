-- name: CreateOrder :one
INSERT INTO orders (
    USER_ID,
    TOTAL_PRICE,
    STATUS,
    SHIPPING_ADDRESS_ID
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetOrderList :many
SELECT 
    * 
FROM 
    orders
WHERE 
    user_id = $1
ORDER BY created_at DESC;

-- name: UpdateOrderAddress :one
UPDATE 
    ORDERS
SET 
    SHIPPING_ADDRESS_ID = $2,
    UPDATED_AT = NOW()
WHERE 
    ORDER_ID = $1 RETURNING *;

-- name: UpdateOrderStatus :one
UPDATE 
    ORDERS
SET 
    status = $2,
    updated_at = NOW()
WHERE 
    ORDER_ID = $1 RETURNING *;