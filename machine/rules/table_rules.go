package rules

import "../../objects"

type TableRules interface {
	//检验条件
	Condition( table objects.Table )
	//进行处理
	Action( table objects.Table )
}

//==================================单轮结算====================================
type SettleForRoundRule struct {}
func (this *SettleForRoundRule) Condition( table objects.Table ) {

}
func (this *SettleForRoundRule) Action( table objects.Table ) {

}

//==================================流局====================================
type LiuJuRule struct {}
func (this *LiuJuRule) Condition( table objects.Table ) {

}
func (this *LiuJuRule) Action( table objects.Table ) {

}