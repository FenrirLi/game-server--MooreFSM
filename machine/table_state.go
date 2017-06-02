package machine

import (
	"log"
	"container/list"
	"reflect"
)

type TableState interface {
	Enter( table *Table )
	Execute( table *Table, event string )
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
	log.Println("    ====TABLE_RULE_READY checking...")
	TableManagerCondition( table, "TABLE_RULE_READY" )

}
func (this *TableReadyState) Execute( table *Table, event string ) {
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

	//开始发牌
	//todo
	table.CardsRest = list.New()
	table.CardsRest.PushFront(21)
	for position, player := range table.PlayerDict {
		player.Cards_in_hand = []int{position,1,2,3,4,5,6,7,8,9}
		table.PlayerDict[position] = player
	}

	//所有用户触发状态
	for position := range table.PlayerDict {
		table.PlayerDict[position].Machine.Trigger( &PlayerDealState{} )
	}

	//该状态下规则检测
	log.Println("    ====TABLE_RULE_DEAL checking...")
	TableManagerCondition( table, "TABLE_RULE_DEAL" )
}
func (this *TableDealState) Execute( table *Table, event string ) {
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
		log.Println(player.Cards_in_hand)
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
func (this *TableStepState) Execute( table *Table, event string ) {
	log.Println("    ====TABLE EXECUTE STEP STATE")
	interface_table_execute( table, event )
}
func (this *TableStepState) Exit( table *Table ) {
	log.Println("===== TABLE EXIT STEP STATE =====")
}
func (this *TableStepState) NextState( table *Table ) {
	log.Println("    ====TABLE NEXT STEP STATE")
}