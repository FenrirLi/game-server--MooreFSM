package machine

import (
	"time"
	"container/list"
	"math/rand"
	"log"
	"reflect"
)


type Table struct {

	//桌子id
	TableId int

	//创建人
	OwnerId string

	//创建时间
	CreateTime time.Time

	//桌子配置
	Config TableConfig

	//玩家
	PlayerDict map[int] *Player

	//状态机
	Machine *TableMachine

	//庄家位置
	DealerSeat int

	//当前行动人位置
	ActiveSeat int

	//当前出牌
	ActiveCard int

	//当前出牌位置
	DiscardSeat int

	//当前执行事件
	Event string

	//剩余卡牌
	CardsRest *list.List

	//获得提示的玩家位置
	PlayerPrompts []int

	//做出行为选择的玩家位置
	PlayerActions []int

	//牌局记录
	Replay list.List

	//胡牌类型：流局，自摸胡，点炮胡，一炮多响
	WinType int

	//当前局数
	CurRound int

	//赢家
	WinnerList []bool

	//输家
	LoserList []bool

	//杠的栈记录
	KongStack bool

	//放杠人位置
	KongTriggerSeat int

	//斗庄
	DouCard int
	DouCard_num int
	IsDouZhuang bool

}

func CreateTable( oid string ) Table {
	rand.Seed(int64(time.Now().Nanosecond()))
	id := rand.Intn(89999)
	id += 10000
	return Table{
		TableId:111111,
		OwnerId: oid,
		CreateTime: time.Now(),
		Config: NewTableConfig(),
		PlayerDict: make(map[int]*Player),
		//Replay: *list.New(),
		CurRound: 0,
		ActiveSeat:-1,
	}
}

func (self *Table) Enter() {
	//TODO

}

func (self *Table) DisMiss() {
	//TODO

}

func (self *Table) InitTable() Table {
	//TODO
	return *self
}

func (self *Table) NextRound() Table {
	//TODO
	return *self
}

func (this *Table) IsAllReady() bool {
	//人未到齐
	if len( this.PlayerDict ) != this.Config.Max_chairs {
		log.Println("人不齐")
		return false
	}
	//有用户未准备
	for _,player := range this.PlayerDict {
		if !reflect.DeepEqual( player.Machine.CurrentState, &PlayerReadyState{} ) {
			log.Println("未准备")
			return false
		}
	}
	return true
}

func (self *Table) CheckAllActed() {
	//没有提示
	if len( self.PlayerPrompts ) == 0 {
		self.KongStack = false
		self.KongTriggerSeat = -1
		//玩家切换到下个状态
		log.Println("table all check next")
		self.PlayerDict[ self.ActiveSeat ].Machine.NextState()
		return
	}

	//还有玩家没有选择操作
	if len( self.PlayerActions ) < len( self.PlayerPrompts ) {
		log.Println("not all checked")
		return
	}
	log.Println("all checked")
	//清除桌子和玩家的提示记录
	self.ClearPrompts()

	//过滤选择出最高权重的操作
	max_weight := 0
	action_seats := []int{}
	for _,v := range self.PlayerActions {
		if player,ok := self.PlayerDict[v]; ok {
			if player.Action.Weight != 0 && player.Action.Weight > max_weight {
				max_weight = player.Action.Weight
			}
		}
	}
	for _,v := range self.PlayerActions {
		if player,ok := self.PlayerDict[v]; ok {
			if player.Action.Weight == max_weight {
				action_seats = append( action_seats, v )
			}
		}
	}

	//选出来的玩家进行操作
	for _,seat := range action_seats {
		if player,ok := self.PlayerDict[seat]; ok {
			action_id := player.Action.ActionId
			if rule,ok := PlayerActionRule[action_id]; ok {
				rule.Action( player )
			}
		}
	}
	//清除桌子和玩家的操作记录
	self.ClearActions()

	//如果是杠牌相关操作，记录"杠"操作记录
	if max_weight == PlayerAction["PLAYER_ACTION_KONG_CONCEALED"] ||
		max_weight == PlayerAction["PLAYER_ACTION_KONG_EXPOSED"] ||
		max_weight == PlayerAction["PLAYER_ACTION_KONG_PONG"] {
		self.KongStack = true
	} else {
		self.KongStack = false
		self.KongTriggerSeat = -1
	}

	//如果是胡牌相关操作
	//自摸
	if max_weight == PlayerAction["PLAYER_ACTION_WIN_DRAW"] {
		if len(action_seats) != 1 {
			log.Fatalln( "Fatal Error : 自摸只可能有一人操作" )
		} else {
			self.WinType = WinTypes["WIN_DRAW"]
		}
		self.Machine.Trigger( &TableEndState{} )
	}
	//点炮胡
	if max_weight == PlayerAction["PLAYER_ACTION_WIN_DISCARD"] {
		if len(action_seats) > 1 {
			self.WinType = WinTypes["WIN_DISCARD_MORE"]
		} else if len( action_seats ) == 1 {
			self.WinType= WinTypes["WIN_DISCARD_ONE"]
		}
		self.Machine.Trigger( &TableEndState{} )
	}


	return
}

//清除当前行为提示玩家记录
func (self *Table) ClearPrompts() {
	self.PlayerPrompts = []int{}
	for _,player := range self.PlayerDict {
		player.DelPrompt()
	}
}

//清除当前行为选择玩家记录
func (self *Table) ClearActions() {
	self.PlayerActions = []int{}
	for _,player := range self.PlayerDict {
		player.DelAction()
	}
}

//过滤行为数组，去除低优先级操作
func (self *Table) FilterActions() {

}