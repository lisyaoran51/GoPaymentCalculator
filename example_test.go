package goPaymentCalculator

import (
	"fmt"
	"testing"
)

func TestEmptyPayment(t *testing.T) {
	var logString string

	member := NewNormalMember(MEMBER_PREVILEDGE_NONE, 1000, 150)
	var costCoin, costPoint float32 = 1000.0, 100.0

	resultPoint, resultcostCoin, _ := NewPaymentCalculator().Calculate(member, costPoint, costCoin)

	logString = fmt.Sprintf("before: point:%f, coin:%f", costPoint, costCoin)
	t.Log(logString)
	logString = fmt.Sprintf("after: point:%f, coin:%f", resultPoint, resultcostCoin)
	t.Log(logString)
}

func TestRedeem(t *testing.T) {
	var logString string

	member := NewNormalMember(MEMBER_PREVILEDGE_NONE, 1000, 150)
	var costCoin, costPoint float32 = 1000.0, 0.0

	resultPoint, resultcostCoin, _ := NewRedeemCalculator(&RedeemCalculatorConfig{
		Ratio: 0.9,
	}).Calculate(member, costPoint, costCoin)

	logString = fmt.Sprintf("before: point:%f, coin:%f", costPoint, costCoin)
	t.Log(logString)
	logString = fmt.Sprintf("after: point:%f, coin:%f", resultPoint, resultcostCoin)
	t.Log(logString)
}

func TestDiscount(t *testing.T) {
	var logString string

	member := NewNormalMember(MEMBER_PREVILEDGE_VIP1, 1000, 150)
	var costCoin, costPoint float32 = 1000.0, 100.0

	resultPoint, resultcostCoin, _ := NewDiscountCalculator(&DiscountCalculatorConfig{
		DiscountMap: map[MEMBER_PREVILEDGE]float32{
			MEMBER_PREVILEDGE_VIP1: 0.9,
			MEMBER_PREVILEDGE_VIP2: 0.8,
			MEMBER_PREVILEDGE_VIP3: 0.7,
		},
	}).Calculate(member, costPoint, costCoin)

	logString = fmt.Sprintf("before: point:%f, coin:%f", costPoint, costCoin)
	t.Log(logString)
	logString = fmt.Sprintf("after: point:%f, coin:%f", resultPoint, resultcostCoin)
	t.Log(logString)
}

func TestDiscountRedeem(t *testing.T) {
	var logString string

	member := NewNormalMember(MEMBER_PREVILEDGE_VIP1, 1000, 150)
	var costCoin, costPoint float32 = 1000.0, 0.0

	// Redeem >> Threshold Dicount
	resultPoint, resultcostCoin, _ := NewRedeemCalculator(&RedeemCalculatorConfig{
		Ratio: 0.9,
	}).CompositeBy(NewThresholdDiscountCalculator(&ThresholdDiscountCalculatorConfig{
		DiscountCalculatorConfig: DiscountCalculatorConfig{
			DiscountMap: map[MEMBER_PREVILEDGE]float32{
				MEMBER_PREVILEDGE_VIP1: 0.9,
				MEMBER_PREVILEDGE_VIP2: 0.9,
				MEMBER_PREVILEDGE_VIP3: 0.9,
			},
		},
		PointThreshold: 100,
	})).Calculate(member, costPoint, costCoin)

	logString = fmt.Sprintf("before: point:%f, coin:%f", costPoint, costCoin)
	t.Log(logString)
	logString = fmt.Sprintf("after: point:%f, coin:%f", resultPoint, resultcostCoin)
	t.Log(logString)
}
