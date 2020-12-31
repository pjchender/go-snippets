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
	Created:  "created",
	Paid:     "paid",
	Canceled: "canceled",
	Shipped:  "shipped",
}

// ToOrderStatus 可以將 string 轉成 OrderStatus，若該 string 不存在則回應錯誤
func ToOrderStatus(orderStatus string) (OrderStatus, error) {
	switch orderStatus {
	case Created.MustString():
		return Created, nil
	case Paid.MustString():
		return Paid, nil
	case Canceled.MustString():
		return Canceled, nil
	case Shipped.MustString():
		return Shipped, nil
	default:
		return 0, fmt.Errorf("invalid orderStatus of %v", orderStatus)
	}
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
	text := orderStatues[w]
	return text
}
