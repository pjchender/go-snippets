package enum

import "fmt"

type OrderStatus int

const (
	Created OrderStatus = iota
	Paid
	Canceled
	Shipped
)

var orderStatues = map[OrderStatus]string{
	Created:  "Created",
	Paid:     "Paid",
	Canceled: "Canceled",
	Shipped:  "Shipped",
}

func (w OrderStatus) String() (string, error) {

	text, ok := orderStatues[w]
	if !ok {
		return "", fmt.Errorf("no available enum of %v", w)
	}

	return text, nil
}
