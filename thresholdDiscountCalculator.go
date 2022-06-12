package goPaymentCalculator


/*
 ********************************************
 *											*
 *											*
 *        ThresholdDiscountCalculator       *
 *											*
 *											*
 ********************************************
 */

type ThresholdDiscountCalculatorConfig struct {
	DiscountCalculatorConfig
	PointThreshold float32
}

type ThresholdDiscountCalculator struct {
	//DiscountCalculator
	// implementing compoistion over inheritance
	DiscountCalculator
	pointThreshold float32
}

func NewThresholdDiscountCalculator(config *ThresholdDiscountCalculatorConfig) *ThresholdDiscountCalculator {
	thresholdDiscountCalculator := &ThresholdDiscountCalculator{
		DiscountCalculator: *NewDiscountCalculator(&config.DiscountCalculatorConfig),
		pointThreshold:     config.PointThreshold,
	}
	thresholdDiscountCalculator.PaymentCalculator = thresholdDiscountCalculator
	return thresholdDiscountCalculator
}

func (d *ThresholdDiscountCalculator) calculate(memeber PaymentMember, costPoint, costCoin float32, callChild bool) (point, coin float32, err error) {
	point, coin, _ = d.EmptyCalculator.Calculate(memeber, costPoint, costCoin)
	if d.child != nil && callChild {
		point, coin, err = d.child.Calculate(memeber, point, coin)
	}

	if point > d.pointThreshold {
		point, coin, err = d.DiscountCalculator.CalculateNonComposite(memeber, point, coin)
	}
	return
}

func (d *ThresholdDiscountCalculator) Calculate(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	return d.calculate(memeber, costPoint, costCoin, true)
}

func (d *ThresholdDiscountCalculator) CalculateNonComposite(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	return d.calculate(memeber, costPoint, costCoin, false)
}
