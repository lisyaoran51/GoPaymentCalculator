package goPaymentCalculator

/*
 ********************************
 *								*
 *								*
 *	     EmptyCalculator	   	*
 *								*
 *								*
 ********************************
 */

type EmptyCalculator struct {
	PaymentCalculator
	child PaymentCalculator
}

func NewEmptyCalculator() *EmptyCalculator {
	emptyCalculator := &EmptyCalculator{}
	emptyCalculator.PaymentCalculator = emptyCalculator
	return emptyCalculator
}

func (e *EmptyCalculator) CompositeBy(parentCalculator PaymentCalculator) PaymentCalculator {
	parentCalculator.AddChild(e)
	return parentCalculator
}

func (e *EmptyCalculator) Calculate(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	if memeber.GetCoin() < costCoin {
		err = ErrorCoinNotEnough
		return
	}
	if memeber.GetPoint() < costPoint {
		err = ErrorPointNotEnough
		return
	}
	coin = costCoin
	point = costPoint
	return
}

func (e *EmptyCalculator) AddChild(childCalculator PaymentCalculator) {
	e.child = childCalculator
}
