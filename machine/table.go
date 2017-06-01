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
	PlayerDict map[int] Player

	//状态机
	Machine *TableMachine

	//当前状态
	Status TableStatus

	//庄家位置
	DealerSeat int

	//当前行动人位置
	ActiveSeat int

	//当前出牌
	ActiveCard int

	//当前出牌位置
	DiscardSeat int

	//event map[int]

	//剩余卡牌
	CardsRest list.List

	//获得提示的玩家位置
	PlayerPrompts map[int] bool

	//做出行为选择的玩家位置
	PlayerActions map[int] bool

	//牌局记录
	Replay list.List

	//胡牌类型：流局，自摸胡，点炮胡，一炮多响
	WinType int

	//当前局数
	CurRound int

	//赢家
	WinnerList map[int] bool

	//输家
	LoserList map[int] bool

	//杠的栈记录
	KongStack bool

	//放杠人位置
	KongTrigger_seat int

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
		PlayerDict: make(map[int] Player),
		//Replay: *list.New(),
		CurRound: 0,
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
		if reflect.DeepEqual( player.Status, PlayerEventStatus["EVENT_READY"] ) {
			log.Println("未准备")
			return false
		}
	}
	return true
}

func (self *Table) IsAllActed() bool {

	return true
}

//清除当前行为提示玩家记录
func (self *Table) ClearPrompts() {
	self.PlayerPrompts = make(map[int] bool)
}

//清除当前行为选择玩家记录
func (self *Table) ClearActions() {
	self.PlayerActions = make(map[int] bool)
}

//过滤行为数组，去除低优先级操作
func (self *Table) FilterActions() {

}