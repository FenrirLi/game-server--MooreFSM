package machine

type PlayerRules interface {
	//检验条件
	Condition( player Player ) bool
	//进行处理
	Action( player Player )
}

//==================================单轮结算====================================
//type SettleForRoundRule struct {}
//func (this *SettleForRoundRule) Condition( table Table ) bool {
//	return false
//}
//func (this *SettleForRoundRule) Action( table Table ) {
//
//}