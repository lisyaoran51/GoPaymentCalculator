package main

import (
	"fmt"

	. "github.com/lisyaoran51/goPaymentCalculator"
)

func main() {
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

	fmt.Printf("before: point:%f, coin:%f\n", costPoint, costCoin)
	fmt.Printf("after: point:%f, coin:%f\n", resultPoint, resultcostCoin)
}
