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
	WinnerList []int

	//输家
	LoserList []int

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
		CurRound: 1,
		ActiveSeat:-1,
	}
}

func (self *Table) Enter() {
	//TODO

}

func (self *Table) DisMiss() {
	//房间进入结算状态，进行相关操作
	self.Machine.Trigger( &TableSettleForRoomState{} )

	//删除房间内用户
	for _,v := range self.PlayerDict {
		delete( GLOBAL_USER, v.Uid )
	}
	//删除桌子
	delete( GLOBAL_TABLE, self.TableId )

}

func (self *Table) InitTable() Table {
	//TODO
	return *self
}

func (self *Table) InitRound() {
	//置为-1，在进入step状态后会取庄家位作为行动人
	self.ActiveSeat = -1
	//当前出牌
	self.ActiveCard = 0
	//当前出牌位置
	self.DiscardSeat = -1
	//获得提示的玩家位置
	self.PlayerPrompts = []int{}
	//做出行为选择的玩家位置
	self.PlayerActions = []int{}
	//牌局记录
	self.Replay = *list.New()
	//胡牌类型：流局，自摸胡，点炮胡，一炮多响
	self.WinType = 0
	//赢家
	self.WinnerList = []int{}
	//输家
	self.LoserList = []int{}
	//杠的栈记录
	self.KongStack = false
	//放杠人位置
	self.KongTriggerSeat = -1
	//玩家重置
	for _,p := range self.PlayerDict {
		p.initRound()
	}
}

func (self *Table) IsAllReady() bool {
	//人未到齐
	if len( self.PlayerDict ) != self.Config.MaxChairs {
		log.Println("人不齐")
		return false
	}
	//有用户未准备
	for _,player := range self.PlayerDict {
		if !reflect.DeepEqual( player.Machine.CurrentState, &PlayerReadyState{} ) {
			log.Println(player.Uid,"未准备")
			return false
		}
	}

	self.Machine.Trigger( &TableReadyState{} )
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
	log.Println("+++操作中+++",max_weight,"   ",action_seats)
	for _,seat := range action_seats {
		if player,ok := self.PlayerDict[seat]; ok {
			action_id := player.Action.ActionId
			log.Println(action_id)
			if rule,ok := PlayerActionRule[action_id]; ok {
				log.Println(reflect.TypeOf(rule).String())
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

func (self *Table) Shuffle() []int {
	vals := []int{
		1,1,1,1,
		2,2,2,2,
		3,3,3,3,
		4,4,4,4,
		5,5,5,5,
		6,6,6,6,
		7,7,7,7,

		11,11,11,11,
		12,12,12,12,
		13,13,13,13,
		14,14,14,14,
		15,15,15,15,
		16,16,16,16,
		17,17,17,17,
		18,18,18,18,
		19,19,19,19,

		21,21,21,21,
		22,22,22,22,
		23,23,23,23,
		24,24,24,24,
		25,25,25,25,
		26,26,26,26,
		27,27,27,27,
		28,28,28,28,
		29,29,29,29,

		31,31,31,31,
		32,32,32,32,
		33,33,33,33,
		34,34,34,34,
		35,35,35,35,
		36,36,36,36,
		37,37,37,37,
		38,38,38,38,
		39,39,39,39,

	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(vals)) {
		val := vals[i]
		log.Println(val)
	}
	return vals
}