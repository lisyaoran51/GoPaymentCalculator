package goPaymentCalculator

type MEMBER_PREVILEDGE int

const (
	MEMBER_PREVILEDGE_NONE MEMBER_PREVILEDGE = iota
	MEMBER_PREVILEDGE_VIP1
	MEMBER_PREVILEDGE_VIP2
	MEMBER_PREVILEDGE_VIP3
)

type PaymentMember interface {
	SetPriviledge(MEMBER_PREVILEDGE)
	GetPriviledge() MEMBER_PREVILEDGE
	SetCoin(float32)
	GetCoin() float32
	SetPoint(float32)
	GetPoint() float32
}
