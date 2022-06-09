package paymentCalculator

/*
 ********************************************
 *											*
 *											*
 * VIPDiscountPointRedeemPaymentCalculator  *
 *											*
 *											*
 ********************************************
 */

type VIPDiscountPointRedeemPaymentCalculatorConfig struct {
	VIPDiscountPaymentCalculatorConfig
	PointThreshold float32
}

type VIPDiscountPointRedeemPaymentCalculator struct {
	VIPDiscountPaymentCalculator
	pointThreshold float32
}

func NewVIPDiscountPointRedeemPaymentCalculator(config *VIPDiscountPointRedeemPaymentCalculatorConfig) *VIPDiscountPointRedeemPaymentCalculator {
	vipDiscountPointRedeemPaymentCalculator := &VIPDiscountPointRedeemPaymentCalculator{
		VIPDiscountPaymentCalculator: *NewVIPDiscountPaymentCalculator(&config.VIPDiscountPaymentCalculatorConfig),
		pointThreshold:               config.PointThreshold,
	}
	vipDiscountPointRedeemPaymentCalculator.PaymentCalculator = vipDiscountPointRedeemPaymentCalculator
	return vipDiscountPointRedeemPaymentCalculator
}

func (p *VIPDiscountPointRedeemPaymentCalculator) Calculate(paymentMemeber PaymentMember, costPoint, costCoin float32) (point, coin float32, err error) {
	point, coin, err = p.child.Calculate(paymentMemeber, costPoint, costCoin)

	if point > p.pointThreshold {
		point, coin, err = p.VIPDiscountPaymentCalculator.Calculate(paymentMemeber, point, coin)
	}
	return
}
