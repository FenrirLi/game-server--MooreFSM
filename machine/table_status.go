package machine

import "log"

type TableStatus interface {
	Enter( table *Table )
	Execute( table *Table, event TableStatus )
	Exit( table *Table )
	NextStatus( table *Table )
}

//===========================TableReadyStatus===========================
type TableReadyStatus struct{}
func (this *TableReadyStatus) Enter( table *Table ) {
	log.Println("-----table enter ready status-----")
	//定位上下家
	for seat,player := range table.PlayerDict {
		//下家
		next_seat := seat + 1;
		if  next_seat >= table.Config.Max_chairs {
			next_seat -= table.Config.Max_chairs
		}
		player.Next_seat = next_seat

		//上家
		prev_seat := seat - 1
		if  prev_seat < 0 {
			prev_seat += table.Config.Max_chairs
		}
		player.Prev_seat = prev_seat
	}
	//该状态下规则检测
	log.Println("=====  TABLE_RULE_READY checking...  =====")
	ManagerCondition( table, "TABLE_RULE_READY" )

}
func (this *TableReadyStatus) Execute( table *Table, event TableStatus ) {
	log.Println("-----table execute ready status-----")
}
func (this *TableReadyStatus) Exit( table *Table ) {
	log.Println("-----table exit ready status-----")
}
func (this *TableReadyStatus) NextStatus( table *Table ) {
	log.Println("-----table next ready status-----")
	table.Machine.Trigger( &TableDealStatus{} )
}


//===========================TableDealStatus===========================
type TableDealStatus struct{}
func (this *TableDealStatus) Enter( table *Table ) {
	log.Println("-----table enter deal status-----")
	//开始发牌
	//todo

	for position, player := range table.PlayerDict {
		player.Cards_in_hand = []int{position,1,2,3,4,5,6,7,8,9}
		table.PlayerDict[position] = player
	}

	//该状态下规则检测
	log.Println("=====  TABLE_RULE_DEAL checking...  =====")
	ManagerCondition( table, "TABLE_RULE_DEAL" )
}
func (this *TableDealStatus) Execute( table *Table, event TableStatus ) {
	log.Println("-----table execute deal status-----")
}
func (this *TableDealStatus) Exit( table *Table ) {
	log.Println("-----table exit deal status-----")
}
func (this *TableDealStatus) NextStatus( table *Table ) {
	log.Println("-----table next deal status-----")
	for _, player := range table.PlayerDict {
		log.Println(player.Cards_in_hand)
	}
}