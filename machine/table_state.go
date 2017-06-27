package machine

import (
	"log"
	"container/list"
	"reflect"
)

type TableState interface {
	Enter( table *Table )
	Execute( table *Table, event string, request_body []byte)
	Exit( table *Table )
	NextState( table *Table )
}

func interface_table_execute( table *Table, event string ) bool {
	//当前状态下执行event事件
	define_event, ok := TableEvent[event]
	if ok && event == define_event {
		log.Printf("    ====TABLE ======= ACTIVE == %s ",event)
	} else {
		log.Println("TableDealState error call event "+event)
	}
	return ok
}

//===========================TableReadyState===========================
type TableReadyState struct{}
func (this *TableReadyState) Enter( table *Table ) {
	log.Println("===== TABLE ENTER READY STATE =====")
	//初始化单轮数据
	table.InitRound()
	//如果是第一局，先将玩家上下家定义好
	if table.CurRound == 1 {
		//定位上下家
		for seat, player := range table.PlayerDict {
			//下家
			next_seat := seat + 1;
			if next_seat >= table.Config.MaxChairs {
				next_seat -= table.Config.MaxChairs
			}
			player.NextSeat = next_seat

			//上家
			prev_seat := seat - 1
			if prev_seat < 0 {
				prev_seat += table.Config.MaxChairs
			}
			player.PrevSeat = prev_seat
		}
	} else {
		//不是第一局，局数后推，上轮数据清空
		table.CurRound++
	}
	//该状态下规则检测
	log.Println("    ====TABLE_RULE_READY checking...")
	TableManagerCondition( table, "TABLE_RULE_READY" )

}
func (this *TableReadyState) Execute( table *Table, event string, request_body []byte) {
	log.Println("    ====TABLE EXECUTE READY STATE")
	interface_table_execute( table, event )
}
func (this *TableReadyState) Exit( table *Table ) {
	log.Println("===== TABLE EXIT READY STATE =====")
}
func (this *TableReadyState) NextState( table *Table ) {
	log.Println("      TABLE NEXT READY STATE")
	table.Machine.Trigger( &TableDealState{} )
}


//===========================TableDealState===========================
type TableDealState struct{}
func (this *TableDealState) Enter( table *Table ) {
	log.Println("===== TABLE ENTER DEAL STATE =====")

	//确定庄家
	table.DealerSeat = 0

	//开始发牌
	table.CardsRest = list.New()
	cards := table.Shuffle()
	for _,v := range cards {
		table.CardsRest.PushFront(v)
	}
	for position, player := range table.PlayerDict {
		for i:=1 ; i <= 13 ; i ++ {
			e := player.Table.CardsRest.Front()
			player.Table.CardsRest.Remove(e)
			draw_card := e.Value.(int)
			player.CardsInHand[i] = draw_card
		}
		player.CardsWin = ReadyHand( player.CardsInHand )
		table.PlayerDict[position] = player
	}

	//所有用户触发状态
	for position := range table.PlayerDict {
		table.PlayerDict[position].Machine.Trigger( &PlayerDealState{} )
	}
}
func (this *TableDealState) Execute( table *Table, event string, request_body []byte) {
	log.Println("    ====TABLE EXECUTE DEAL STATE")
	if interface_table_execute( table, event ) {
		//检测是否所有玩家都已进入等待状态
		for _,player := range table.PlayerDict {
			//如果有玩家不在等待状态
			if !reflect.DeepEqual( player.Machine.CurrentState, &PlayerWaitState{} ) {
				log.Printf("     TABLE EXECUTE DEAL : PLAYER %s NOT IN WAIT STATE",player.Uid)
				return
			}
		}
		log.Println("    ====TABLE EXECUTE DEAL : ALL IN WAIT STATE")
		table.Machine.Trigger( &TableStepState{} )
	}
}
func (this *TableDealState) Exit( table *Table ) {
	log.Println("===== TABLE EXIT DEAL STATE =====")
}
func (this *TableDealState) NextState( table *Table ) {
	log.Println("    ====TABLE NEXT DEAL STATE")
	for _, player := range table.PlayerDict {
		log.Println(player.CardsInHand)
	}
	table.Machine.Trigger( &TableStepState{} )
}


//===========================TableStepState===========================
type TableStepState struct{}
func (this *TableStepState) Enter( table *Table ) {
	log.Println("===== TABLE ENTER STEP STATE =====")
	//该状态下规则检测
	log.Println("    ====TABLE_RULE_STEP checking...")
	TableManagerCondition( table, "TABLE_RULE_STEP" )
}
func (this *TableStepState) Execute( table *Table, event string, request_body []byte) {
	log.Println("    ====TABLE EXECUTE STEP STATE")
	interface_table_execute( table, event )
}
func (this *TableStepState) Exit( table *Table ) {
	log.Println("===== TABLE EXIT STEP STATE =====")
}
func (this *TableStepState) NextState( table *Table ) {
	log.Println("    ====TABLE NEXT STEP STATE")
	before_id := table.ActiveSeat
	//确定当前步骤执行人
	if table.ActiveSeat >= 0 {
		table.ActiveSeat = table.PlayerDict[ table.ActiveSeat ].NextSeat
	} else {
		table.ActiveSeat = table.DealerSeat
	}
	active_player := table.PlayerDict[ table.ActiveSeat ]
	log.Println("-------------------",before_id,"-----",table.ActiveSeat,"-----------")
	//桌子切换状态
	table.Machine.Trigger( &TableWaitState{} )
	active_player.Machine.Trigger( &PlayerDrawState{} )

}


//===========================TableWaitState===========================
type TableWaitState struct{}
func (this *TableWaitState) Enter( table *Table ) {
	log.Println("===== TABLE ENTER WAIT STATE =====")
}
func (this *TableWaitState) Execute( table *Table, event string, request_body []byte) {
	log.Println("    ====TABLE EXECUTE WAIT STATE")
	if interface_table_execute( table, event ) {
		switch event {
			//执行步骤
			case TableEvent["TABLE_EVENT_STEP"] :
				table.Machine.Trigger( &TableStepState{} )
			//结束
			case TableEvent["TABLE_EVENT_END"] :
				table.Machine.Trigger( &TableEndState{} )

			default:
				log.Println("----- no such event ",event," ----- ")
		}
	}
}
func (this *TableWaitState) Exit( table *Table ) {
	log.Println("===== TABLE EXIT WAIT STATE =====")
}
func (this *TableWaitState) NextState( table *Table ) {
	log.Println("    ====TABLE NEXT WAIT STATE")
}


//===========================TableEndState===========================
type TableEndState struct{}
func (this *TableEndState) Enter( table *Table ) {
	log.Println("===== TABLE ENTER END STATE =====")
	log.Println("    ====TABLE_RULE_END checking...")
	TableManagerCondition( table, "TABLE_RULE_END" )
}
func (this *TableEndState) Execute( table *Table, event string, request_body []byte) {
	log.Println("    ====TABLE EXECUTE END STATE")
	interface_table_execute( table, event )
}
func (this *TableEndState) Exit( table *Table ) {
	log.Println("===== TABLE EXIT END STATE =====")
}
func (this *TableEndState) NextState( table *Table ) {
	log.Println("    ====TABLE NEXT END STATE")
	table.Machine.Trigger( &TableSettleForRoundState{} )
}


//===========================TableSettleForRoundState===========================
type TableSettleForRoundState struct{}
func (this *TableSettleForRoundState) Enter( table *Table ) {
	log.Println("===== TABLE ENTER SETTLE FOR ROUND STATE =====")

	//玩家总局数据累计
	for _, player := range table.PlayerDict {
		player.ScoreTotal += player.Score
		player.ScoreKongTotal += player.KongScore
		player.KongExposedTotal += player.KongExposedCnt
		player.KongPongTotal += player.KongPongCnt
		player.KongConcealedTotal += player.KongConcealedCnt
		player.KongDiscardTotal += player.KongDiscardCnt
	}

	log.Println("    ====TABLE_RULE_SETTLE_FOR_ROUND checking...")
	TableManagerCondition( table, "TABLE_RULE_SETTLE_FOR_ROUND" )
}
func (this *TableSettleForRoundState) Execute( table *Table, event string, request_body []byte) {
	log.Println("    ====TABLE EXECUTE SETTLE FOR ROUND STATE")
	interface_table_execute( table, event )
}
func (this *TableSettleForRoundState) Exit( table *Table ) {
	log.Println("===== TABLE EXIT SETTLE FOR ROUND STATE =====")
}
func (this *TableSettleForRoundState) NextState( table *Table ) {
	log.Println("    ====TABLE NEXT SETTLE FOR ROUND STATE")
	table.Machine.Trigger( &TableRestartState{} )
}


//===========================TableSettleForRoomState===========================
type TableSettleForRoomState struct{}
func (this *TableSettleForRoomState) Enter( table *Table ) {
	log.Println("===== TABLE ENTER SETTLE FOR ROOM STATE =====")
	//广播
	//TODO
}
func (this *TableSettleForRoomState) Execute( table *Table, event string, request_body []byte) {
	log.Println("    ====TABLE EXECUTE SETTLE FOR ROOM STATE")
	interface_table_execute( table, event )
}
func (this *TableSettleForRoomState) Exit( table *Table ) {
	log.Println("===== TABLE EXIT SETTLE FOR ROOM STATE =====")
}
func (this *TableSettleForRoomState) NextState( table *Table ) {
	log.Println("    ====TABLE NEXT SETTLE FOR ROOM STATE")
}


//===========================TableRestartState===========================
type TableRestartState struct{}
func (this *TableRestartState) Enter( table *Table ) {
	log.Println("===== TABLE ENTER RESTART STATE =====")
	for _,p := range table.PlayerDict{
		log.Println(p)
	}
}
func (this *TableRestartState) Execute( table *Table, event string, request_body []byte) {
	log.Println("    ====TABLE EXECUTE RESTART STATE")
	interface_table_execute( table, event )
}
func (this *TableRestartState) Exit( table *Table ) {
	log.Println("===== TABLE EXIT RESTART STATE =====")
}
func (this *TableRestartState) NextState( table *Table ) {
	log.Println("    ====TABLE NEXT RESTART STATE")
}