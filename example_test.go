package paymentCalculator

import (
	"testing"
)

func TestEmptyPaymentCalculator(t *testing.T) {

	paymentMember := NewNormalPaymentMember(MEMBER_PREVILEDGE_NONE, 1000, 150)
	var costPoint, costCoin float32 = 100.0, 1000.0

	resultPoint, resultcostCoin, _ := NewPaymentCalculator().Calculate(paymentMember, costPoint, costCoin)

	t.Log("before: point-%d, coin-%d\n", costPoint, costCoin)
	t.Log("after: point-%d, coin-%d\n", resultPoint, resultcostCoin)
}

func TestEmptyPaymentCalculator(t *testing.T) {

	paymentMember := NewNormalPaymentMember(MEMBER_PREVILEDGE_NONE, 1000, 150)
	var costPoint, costCoin float32 = 100.0, 1000.0

	resultPoint, resultcostCoin, _ := NewPaymentCalculator().Calculate(paymentMember, costPoint, costCoin)

	t.Log("before: point-%d, coin-%d\n", costPoint, costCoin)
	t.Log("after: point-%d, coin-%d\n", resultPoint, resultcostCoin)
}
