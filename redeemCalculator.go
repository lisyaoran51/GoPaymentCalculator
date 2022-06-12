package goPaymentCalculator

import (
	"math"
)

/*
 ********************************
 *								*
 *								*
 *       RedeemCalculator       *
 *								*
 *								*
 ********************************
 */

type RedeemCalculatorConfig struct {
	//Ratio
	// the coin:point ratio
	Ratio float32
}

type RedeemCalculator struct {
	//EmptyCalculator
	// implementing compoistion over inheritance
	EmptyCalculator
	redeemRatio float32
}

func NewRedeemCalculator(config *RedeemCalculatorConfig) *RedeemCalculator {
	redeemCalculator := &RedeemCalculator{
		EmptyCalculator: *NewEmptyCalculator(),
		redeemRatio:     config.Ratio,
	}
	redeemCalculator.PaymentCalculator = redeemCalculator
	return redeemCalculator
}

func (r *RedeemCalculator) calculate(memeber PaymentMember, costPoint, costCoin float32, callChild bool) (point, coin float32, err error) {
	point, coin, _ = r.EmptyCalculator.Calculate(memeber, costPoint, costCoin)
	if r.child != nil && callChild {
		point, coin, err = r.child.Calculate(memeber, point, coin)
	}

	tempPoint := float32(math.Min(float64(coin*r.redeemRatio+point), float64(memeber.GetPoint())))
	coin -= (tempPoint - point) / r.redeemRatio
	point = tempPoint
	return
}

func (r *RedeemCalculator) Calculate(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	return r.calculate(memeber, costPoint, costCoin, true)
}

func (r *RedeemCalculator) CalculateNonComposite(memeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	return r.calculate(memeber, costPoint, costCoin, false)
}
