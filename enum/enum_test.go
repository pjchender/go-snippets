package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnum(t *testing.T) {
	assertWithT := assert.New(t)
	var receiveData int

	/* Canceled */
	receiveData = 0
	orderStatus, err := OrderStatus(receiveData).String()
	assertWithT.NoError(err)
	assertWithT.Equal(orderStatues[Created], orderStatus)

	/* Paid */
	receiveData = 1
	orderStatus, err = OrderStatus(receiveData).String()
	assertWithT.NoError(err)
	assertWithT.Equal(orderStatues[Paid], orderStatus)

	/* Error */
	receiveData = 10 // out of range
	orderStatus, err = OrderStatus(receiveData).String()
	assertWithT.Error(err)
	assertWithT.Equal("", orderStatus)
}
