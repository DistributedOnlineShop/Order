package util

import (
	"math/big"
	"math/rand/v2"

	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateRandomStatus() string {
	var orderStatuses = []string{
		"PENDING",
		"PROCESSING",
		"AWAITING_SHIPMENT",
		"SHIPPED",
		"OUT_FOR_DELIVERY",
		"DELIVERED",
		"COMPLETED",
		"CANCELLED",
		"RETURN_REQUESTED",
		"RETURNED",
	}

	return orderStatuses[rand.IntN(len(orderStatuses))]
}

func GenerateRandomNumeric() pgtype.Numeric {
	intPart := rand.IntN(100000)
	fracPart := rand.IntN(100)
	value := int64(intPart) + int64(fracPart)/100.0

	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   -2,
		Valid: true,
	}
}
