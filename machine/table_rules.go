package machine

import "log"

type TableRules interface {
	//检验条件
	Condition( table *Table ) bool
	//进行处理
	Action( table *Table )
}

//==================================单轮结算====================================
type TableSettleForRoundRule struct {}
func (self *TableSettleForRoundRule) Condition( table *Table ) bool {
	if table.CurRound > table.Config.MaxRounds {
		return true
	} else {
		return false
	}
}
func (self *TableSettleForRoundRule) Action( table *Table ) {
	
}

//==================================流局====================================
type TableLiuJuRule struct {}
func (self *TableLiuJuRule) Condition( table *Table ) bool{
	if table.CardsRest.Len() == 0 {
		log.Println("=====rule liuju true=====")
		return true
	} else {
		log.Println("=====rule liuju false=====")
		return false
	}
}
func (self *TableLiuJuRule) Action( table *Table ) {
	log.Println("=====rule liuju action=====")
	table.Machine.Trigger( &TableEndState{} )
}