package util

import (
	"fmt"
	"math/big"
	"math/rand/v2"

	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateStatus() string {
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

func GenerateNumeric() pgtype.Numeric {
	intPart := rand.IntN(100000)
	fracPart := rand.IntN(100)
	value := int64(intPart) + int64(fracPart)/100.0

	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   -2,
		Valid: true,
	}
}

func GenerateProductID() string {
	return "P" + fmt.Sprintf("%07d", rand.IntN(9999999)+1)
}

func GeneratePvID() string {
	return "PV" + fmt.Sprintf("%07d", rand.IntN(999999)+1)
}

func GenerateInt32() int32 {
	return rand.Int32N(1000) + 1
}
