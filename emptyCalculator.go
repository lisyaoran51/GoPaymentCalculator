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
	PaymentCalculatorImpl
}

func NewPaymentCalculator() PaymentCalculator {
	return NewEmptyCalculator()
}

func NewEmptyCalculator() *EmptyCalculator {
	emptyCalculator := &EmptyCalculator{
		PaymentCalculatorImpl: *NewPaymentCalculatorImpl(),
	}
	emptyCalculator.PaymentCalculator = emptyCalculator
	return emptyCalculator
}

func (e *EmptyCalculator) calculate(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {

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

func (e *EmptyCalculator) Calculate(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	return e.calculate(memeber, costPoint, costCoin)
}

func (e *EmptyCalculator) CalculateNonComposite(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	return e.calculate(memeber, costPoint, costCoin)
}
