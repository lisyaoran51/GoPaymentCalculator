package goPaymentCalculator

import "fmt"

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

func (d *ThresholdDiscountCalculator) Calculate(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	point, coin, _ = d.EmptyCalculator.Calculate(memeber, costPoint, costCoin)
	if d.child != nil {
		point, coin, err = d.child.Calculate(memeber, point, coin)
		fmt.Printf("ThresholdDiscountCalculator after: point:%f, coin:%f\n", point, coin)
	}

	if point > d.pointThreshold {
		point, coin, err = d.DiscountCalculator.Calculate(memeber, point, coin)
	}
	return
}
