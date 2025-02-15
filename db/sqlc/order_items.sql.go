// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: order_items.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO order_items (
    OI_ID,
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
    $6,
    $7
) RETURNING oi_id, order_id, product_id, pv_id, quantity, price, total, created_at, updated_at
`

type CreateOrderItemParams struct {
	OiID      uuid.UUID      `json:"oi_id"`
	OrderID   uuid.UUID      `json:"order_id"`
	ProductID string         `json:"product_id"`
	PvID      pgtype.Text    `json:"pv_id"`
	Quantity  int32          `json:"quantity"`
	Price     pgtype.Numeric `json:"price"`
	Total     pgtype.Numeric `json:"total"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRow(ctx, createOrderItem,
		arg.OiID,
		arg.OrderID,
		arg.ProductID,
		arg.PvID,
		arg.Quantity,
		arg.Price,
		arg.Total,
	)
	var i OrderItem
	err := row.Scan(
		&i.OiID,
		&i.OrderID,
		&i.ProductID,
		&i.PvID,
		&i.Quantity,
		&i.Price,
		&i.Total,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOrderItem = `-- name: DeleteOrderItem :exec
DELETE FROM 
    order_items 
WHERE 
    OI_ID = $1
`

func (q *Queries) DeleteOrderItem(ctx context.Context, oiID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteOrderItem, oiID)
	return err
}

const getOrderItemList = `-- name: GetOrderItemList :many
SELECT
    oi_id, order_id, product_id, pv_id, quantity, price, total, created_at, updated_at
FROM
    order_items
`

func (q *Queries) GetOrderItemList(ctx context.Context) ([]OrderItem, error) {
	rows, err := q.db.Query(ctx, getOrderItemList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderItem{}
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.OiID,
			&i.OrderID,
			&i.ProductID,
			&i.PvID,
			&i.Quantity,
			&i.Price,
			&i.Total,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderItemsByOrderId = `-- name: GetOrderItemsByOrderId :one
SELECT 
    oi_id, order_id, product_id, pv_id, quantity, price, total, created_at, updated_at 
FROM 
    order_items 
WHERE 
    ORDER_ID = $1
`

func (q *Queries) GetOrderItemsByOrderId(ctx context.Context, orderID uuid.UUID) (OrderItem, error) {
	row := q.db.QueryRow(ctx, getOrderItemsByOrderId, orderID)
	var i OrderItem
	err := row.Scan(
		&i.OiID,
		&i.OrderID,
		&i.ProductID,
		&i.PvID,
		&i.Quantity,
		&i.Price,
		&i.Total,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateOrderItem = `-- name: UpdateOrderItem :one
UPDATE 
    order_items 
SET 
    QUANTITY = COALESCE($2,QUANTITY),
    PRICE = COALESCE($3,PRICE),
    TOTAL = COALESCE($4,TOTAL),
    UPDATED_AT = NOW()
WHERE 
    OI_ID = $1 RETURNING oi_id, order_id, product_id, pv_id, quantity, price, total, created_at, updated_at
`

type UpdateOrderItemParams struct {
	OiID     uuid.UUID      `json:"oi_id"`
	Quantity int32          `json:"quantity"`
	Price    pgtype.Numeric `json:"price"`
	Total    pgtype.Numeric `json:"total"`
}

func (q *Queries) UpdateOrderItem(ctx context.Context, arg UpdateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRow(ctx, updateOrderItem,
		arg.OiID,
		arg.Quantity,
		arg.Price,
		arg.Total,
	)
	var i OrderItem
	err := row.Scan(
		&i.OiID,
		&i.OrderID,
		&i.ProductID,
		&i.PvID,
		&i.Quantity,
		&i.Price,
		&i.Total,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
