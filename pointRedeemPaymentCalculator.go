package paymentCalculator

import "math"

/*
 ********************************
 *								*
 *								*
 * PointRedeemPaymentCalculator *
 *								*
 *								*
 ********************************
 */

type PointRedeemPaymentCalculatorConfig struct {
	//RedeemRatio
	// the coin:point ratio
	Ratio float32
}

type PointRedeemPaymentCalculator struct {
	EmptyPaymentCalculator
	redeemRatio float32
}

func NewPointRedeemPaymentCalculator(config *PointRedeemPaymentCalculatorConfig) *PointRedeemPaymentCalculator {
	pointRedeemPaymentCalculator := &PointRedeemPaymentCalculator{
		EmptyPaymentCalculator: *NewEmptyPaymentCalculator(),
		redeemRatio:            config.Ratio,
	}
	pointRedeemPaymentCalculator.PaymentCalculator = pointRedeemPaymentCalculator
	return pointRedeemPaymentCalculator
}

func (p *PointRedeemPaymentCalculator) Calculate(paymentMemeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	point, coin, err = p.child.Calculate(paymentMemeber, costPoint, costCoin)

	tempPoint := float32(math.Max(float64(coin*p.redeemRatio+point), float64(paymentMemeber.GetPoint())))
	coin = (tempPoint - point) / p.redeemRatio
	point = tempPoint
	return
}
