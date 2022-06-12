package goPaymentCalculator

type NormalMember struct {
	priviledge MEMBER_PREVILEDGE
	coin       float32
	point      float32
}

func NewNormalMember(priviledge MEMBER_PREVILEDGE, coin float32, point float32) *NormalMember {
	return &NormalMember{
		priviledge: priviledge,
		coin:       coin,
		point:      point,
	}
}

func (n *NormalMember) SetPriviledge(priviledge MEMBER_PREVILEDGE) {
	n.priviledge = priviledge
}

func (n *NormalMember) GetPriviledge() MEMBER_PREVILEDGE {
	return n.priviledge
}

func (n *NormalMember) SetCoin(coin float32) {
	n.coin = coin
}

func (n *NormalMember) GetCoin() float32 {
	return n.coin
}

func (n *NormalMember) SetPoint(point float32) {
	n.point = point
}

func (n *NormalMember) GetPoint() float32 {
	return n.point
}
