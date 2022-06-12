package goPaymentCalculator

import (
	"errors"
)

var (
	ErrorCoinNotEnough  = errors.New("Coin Not Enough")
	ErrorPointNotEnough = errors.New("Point Not Enough")
	ErrorNoPriviledge   = errors.New("No priviledge")
)

type PaymentCalculator interface {

	//CompositeBy
	// the input calculator would be parent of this calculator. then return the input parent
	// calculate flow would be: parent -> child
	CompositeBy(PaymentCalculator) PaymentCalculator
	Calculate(PaymentMember, float32, float32) (point, coin float32, err error)
	AddChild(PaymentCalculator)
}

func NewPaymentCalculator() PaymentCalculator {
	return NewEmptyCalculator()
}
