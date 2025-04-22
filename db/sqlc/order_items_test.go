package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"Order/util"
)

func CreateRandomOrderItem(t *testing.T, orderID uuid.UUID) OrderItem {
	data := CreateOrderItemParams{
		OiID:      util.CreateUUID(),
		OrderID:   orderID,
		ProductID: util.GenerateProductID(),
		PvID:      util.GeneratePvID(),
		Quantity:  util.GenerateInt32(),
		Price:     util.GenerateNumeric(),
		Total:     util.GenerateNumeric(),
	}

	oI, err := testStore.CreateOrderItem(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, oI)
	require.Equal(t, oI.OiID, data.OiID)
	require.Equal(t, oI.OrderID, data.OrderID)
	require.Equal(t, oI.ProductID, data.ProductID)
	require.Equal(t, oI.PvID, data.PvID)
	require.Equal(t, oI.Quantity, data.Quantity)
	require.Equal(t, oI.Price, data.Price)
	require.Equal(t, oI.Total, data.Total)
	require.NotZero(t, oI.CreatedAt)

	return oI
}

func TestCreateOrderItem(t *testing.T) {
	orderID := util.CreateUUID()
	CreateRandomOrderItem(t, orderID)
}

func TestGetOrderItemList(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := util.CreateUUID()
		orderID := CreateRandomOrder(t, id)
		CreateRandomOrderItem(t, orderID.OrderID)
	}

	orderItems, err := testStore.GetOrderItemList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, orderItems)
	require.GreaterOrEqual(t, len(orderItems), 10)
}
func TestGetOrderItemByOrderId(t *testing.T) {
	id := util.CreateUUID()
	order := CreateRandomOrder(t, id)
	for i := 0; i < 10; i++ {
		CreateRandomOrderItem(t, order.OrderID)
	}

	orderItems, err := testStore.GetOrderItemsByOrderId(context.Background(), order.OrderID)
	require.NoError(t, err)
	require.NotEmpty(t, orderItems)
	require.Equal(t, orderItems.OrderID, order.OrderID)
}

func TestUpdateOrderItem(t *testing.T) {
	id := util.CreateUUID()
	order := CreateRandomOrder(t, id)
	oi := CreateRandomOrderItem(t, order.OrderID)

	newData := UpdateOrderItemParams{
		OiID:     oi.OiID,
		Quantity: util.GenerateInt32(),
		Price:    util.GenerateNumeric(),
		Total:    util.GenerateNumeric(),
	}

	updatedOI, err := testStore.UpdateOrderItem(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, updatedOI)
	require.Equal(t, oi.OiID, updatedOI.OiID)
	require.NotEqual(t, oi.Quantity, updatedOI.Quantity)
	require.NotEqual(t, oi.Total, updatedOI.Total)
	require.NotZero(t, updatedOI.UpdatedAt)
}

func TestDeleteOrderItem(t *testing.T) {
	id := util.CreateUUID()
	order := CreateRandomOrder(t, id)
	oi := CreateRandomOrderItem(t, order.OrderID)

	err := testStore.DeleteOrderItem(context.Background(), oi.OiID)
	require.NoError(t, err)
}
