package enum

import "fmt"

type OrderStatus string

const (
	Created  OrderStatus = "created"
	Paid                 = "paid"
	Canceled             = "canceled"
	Shipped              = "shipped"
)

var orderStatues = map[OrderStatus]string{
	Created:  "created",
	Paid:     "paid",
	Canceled: "canceled",
	Shipped:  "shipped",
}

func IsValidOrderStatus(orderStatus string) error {
	_, ok := orderStatues[OrderStatus(orderStatus)]
	if !ok {
		return fmt.Errorf("no available enum of %v", orderStatus)
	}

	return nil
}

func (w OrderStatus) String() (string, error) {
	text, ok := orderStatues[w]
	if !ok {
		return "", fmt.Errorf("no available enum of %v", w)
	}

	return text, nil
}

// MustString will always return string else panic
func (w OrderStatus) MustString() string {
	text, ok := orderStatues[w]
	if !ok {
		panic(fmt.Errorf("no available enum of %v", w))
	}

	return text
}
