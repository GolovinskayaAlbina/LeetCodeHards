package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_IsOrderPrinter(t *testing.T) {
	p := NewOrderPrinter()
	p.Third("3")
	time.Sleep(100)
	p.Second("2")
	time.Sleep(100)
	p.First("1")

	assert.Equal(t, "123", p.GetResult())
}
