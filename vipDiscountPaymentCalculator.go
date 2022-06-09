package paymentCalculator

/*
 ********************************
 *								*
 *								*
 * VIPDiscountPaymentCalculator *
 *								*
 *								*
 ********************************
 */

type VIPDiscountPaymentCalculatorConfig struct {
	DiscountMap map[MEMBER_PREVILEDGE]float32
}

type VIPDiscountPaymentCalculator struct {
	EmptyPaymentCalculator
	discountMap map[MEMBER_PREVILEDGE]float32
}

func NewVIPDiscountPaymentCalculator(config *VIPDiscountPaymentCalculatorConfig) *VIPDiscountPaymentCalculator {
	vipDiscountPaymentCalculator := &VIPDiscountPaymentCalculator{
		EmptyPaymentCalculator: *NewEmptyPaymentCalculator(),
		discountMap:            config.DiscountMap,
	}
	vipDiscountPaymentCalculator.PaymentCalculator = vipDiscountPaymentCalculator
	return vipDiscountPaymentCalculator
}

func (v *VIPDiscountPaymentCalculator) Calculate(paymentMemeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	point, coin, err = v.child.Calculate(paymentMemeber, costPoint, costCoin)
	if discount, ok := v.discountMap[paymentMemeber.GetPriviledge()]; ok {
		point *= discount
		coin *= discount
	}
	return
}
