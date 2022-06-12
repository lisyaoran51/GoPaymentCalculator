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
	// calculate flow would be: child -> parent
	CompositeBy(PaymentCalculator) PaymentCalculator
	Calculate(PaymentMember, float32, float32) (point, coin float32, err error)

	//CalculateNonComposite
	// dosen't call composite child to calculate
	CalculateNonComposite(PaymentMember, float32, float32) (point, coin float32, err error)
	AddChild(PaymentCalculator)
}

func NewPaymentCalculatorImpl() *PaymentCalculatorImpl {
	return &PaymentCalculatorImpl{}
}

type PaymentCalculatorImpl struct {
	PaymentCalculator
	child PaymentCalculator
}

func (p *PaymentCalculatorImpl) CompositeBy(parentCalculator PaymentCalculator) PaymentCalculator {
	parentCalculator.AddChild(p)
	return parentCalculator
}

func (p *PaymentCalculatorImpl) Calculate(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	return p.PaymentCalculator.Calculate(memeber, costPoint, costCoin)
}

func (p *PaymentCalculatorImpl) AddChild(childCalculator PaymentCalculator) {
	p.child = childCalculator
}
