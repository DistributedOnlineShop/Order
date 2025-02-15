package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"Order/util"
)

func CreateRandomOrder(t *testing.T, userID uuid.UUID) Order {
	data := CreateOrderParams{
		OrderID:           util.CreateUUID(),
		UserID:            userID,
		TotalPrice:        util.GenerateRandomNumeric(),
		Status:            util.GenerateRandomStatus(),
		ShippingAddressID: util.CreateUUID(),
	}

	order, err := testStore.CreateOrder(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, order)
	require.Equal(t, order.OrderID, data.OrderID)
	require.Equal(t, order.UserID, data.UserID)
	require.Equal(t, order.TotalPrice, data.TotalPrice)
	require.Equal(t, order.Status, data.Status)
	require.Equal(t, order.ShippingAddressID, data.ShippingAddressID)

	return order
}

func TestCreateOrder(t *testing.T) {
	id := util.CreateUUID()
	CreateRandomOrder(t, id)
}

func TestGetOrderList(t *testing.T) {
	id := util.CreateUUID()
	for i := 0; i < 10; i++ {
		CreateRandomOrder(t, id)
	}

	orders, err := testStore.GetOrderList(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, orders)
	require.GreaterOrEqual(t, len(orders), 10)
}

func TestUpdateOrderAddress(t *testing.T) {
	id := util.CreateUUID()
	order := CreateRandomOrder(t, id)

	newData := UpdateOrderAddressParams{
		OrderID:           order.OrderID,
		ShippingAddressID: util.CreateUUID(),
	}

	updatedOrder, err := testStore.UpdateOrderAddress(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, updatedOrder)
	require.Equal(t, updatedOrder.OrderID, order.OrderID)
	require.NotEqual(t, updatedOrder.ShippingAddressID, order.ShippingAddressID)
	require.NotZero(t, updatedOrder.UpdatedAt)
}

func TestUpdateOrderStatus(t *testing.T) {
	id := util.CreateUUID()
	order := CreateRandomOrder(t, id)

	newData := UpdateOrderStatusParams{
		OrderID: order.OrderID,
		Status:  util.GenerateRandomStatus(),
	}

	updateStatus, err := testStore.UpdateOrderStatus(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, updateStatus)
	require.Equal(t, updateStatus.OrderID, order.OrderID)
	require.NotEqual(t, updateStatus.Status, order.Status)
	require.NotZero(t, updateStatus.UpdatedAt)
}
