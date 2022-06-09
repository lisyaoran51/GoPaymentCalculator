package paymentCalculator

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

type NormalPaymentMember struct {
	priviledge MEMBER_PREVILEDGE
	coin       float32
	point      float32
}

func NewNormalPaymentMember(priviledge MEMBER_PREVILEDGE, coin float32, point float32) *NormalPaymentMember {
	return &NormalPaymentMember{
		priviledge: priviledge,
		coin:       coin,
		point:      point,
	}
}

func (n *NormalPaymentMember) SetPriviledge(priviledge MEMBER_PREVILEDGE) {
	n.priviledge = priviledge
}

func (n *NormalPaymentMember) GetPriviledge() MEMBER_PREVILEDGE {
	return n.priviledge
}

func (n *NormalPaymentMember) SetCoin(coin float32) {
	n.coin = coin
}

func (n *NormalPaymentMember) GetCoin() float32 {
	return n.coin
}

func (n *NormalPaymentMember) SetPoint(point float32) {
	n.point = point
}

func (n *NormalPaymentMember) GetPoint() float32 {
	return n.point
}
