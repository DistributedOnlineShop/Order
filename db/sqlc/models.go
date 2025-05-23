// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Order struct {
	OrderID           uuid.UUID        `json:"order_id"`
	UserID            uuid.UUID        `json:"user_id"`
	TotalPrice        pgtype.Numeric   `json:"total_price"`
	Status            string           `json:"status"`
	ShippingAddressID uuid.UUID        `json:"shipping_address_id"`
	CreatedAt         pgtype.Timestamp `json:"created_at"`
	UpdatedAt         pgtype.Timestamp `json:"updated_at"`
}

type OrderItem struct {
	OiID      uuid.UUID        `json:"oi_id"`
	OrderID   uuid.UUID        `json:"order_id"`
	ProductID string           `json:"product_id"`
	PvID      string           `json:"pv_id"`
	Quantity  int32            `json:"quantity"`
	Price     pgtype.Numeric   `json:"price"`
	Total     pgtype.Numeric   `json:"total"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
