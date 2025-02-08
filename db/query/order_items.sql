-- name: CreateOrderItem :one
INSERT INTO order_items (
    ORDER_ID,
    PRODUCT_ID,
    pv_id,
    QUANTITY,
    PRICE,
    TOTAL
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: GetOrderItemList :many
SELECT
    *
FROM
    order_items;

-- name: GetOrderItemsByOrderId :one
SELECT 
    * 
FROM 
    order_items 
WHERE 
    ORDER_ID = $1;

-- name: UpdateOrderItem :one
UPDATE 
    order_items 
SET 
    QUANTITY = COALESCE($2,QUANTITY),
    PRICE = COALESCE($3,PRICE),
    TOTAL = COALESCE($4,TOTAL),
    UPDATED_AT = NOW()
WHERE 
    OI_ID = $1 RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM 
    order_items 
WHERE 
    OI_ID = $1;