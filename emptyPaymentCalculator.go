package paymentCalculator

/*
 ********************************
 *								*
 *								*
 *	  EmptyPaymentCalculator	*
 *								*
 *								*
 ********************************
 */

type EmptyPaymentCalculator struct {
	PaymentCalculator
	child PaymentCalculator
}

func NewEmptyPaymentCalculator() *EmptyPaymentCalculator {
	emptyPaymentCalculator := &EmptyPaymentCalculator{}
	emptyPaymentCalculator.PaymentCalculator = emptyPaymentCalculator
	return emptyPaymentCalculator
}

func (e *EmptyPaymentCalculator) CompositeBy(parentCalculator PaymentCalculator) PaymentCalculator {
	parentCalculator.AddChild(e)
	return parentCalculator
}

func (e *EmptyPaymentCalculator) Calculate(paymentMember PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	if paymentMember.GetCoin() < costCoin {
		err = ErrorCoinNotEnough
		return
	}
	if paymentMember.GetPoint() < costPoint {
		err = ErrorPointNotEnough
		return
	}
	coin = costCoin
	point = costPoint
	return
}

func (e *EmptyPaymentCalculator) AddChild(childCalculator PaymentCalculator) {
	e.child = childCalculator
}
