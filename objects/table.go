package objects

import (
	"time"
	"container/list"
	"../machine"
	"../define"
)


type Table struct {

	//桌子id
	Table_id string

	//创建人
	Owner string

	//创建时间
	Create_time time.Time

	//桌子配置
	Config TableConfig

	//玩家
	Player_dict map[int] Player

	//状态机
	Machine machine.Machine
	State string

	//庄家位置
	Dealer_seat int

	//当前行动人位置
	Active_seat int

	//当前出牌
	Active_card int

	//当前出牌位置
	Discard_seat int

	//event map[int]

	//剩余卡牌
	Cards_rest list.List

	//获得提示的玩家位置
	Player_prompts map[int] bool

	//做出行为选择的玩家位置
	Player_actions map[int] bool

	//牌局记录
	Replay list.List

	//胡牌类型：流局，自摸胡，点炮胡，一炮多响
	Win_type int

	//当前局数
	Cur_round int

	//赢家
	Winner_list map[int] bool

	//输家
	Loser_list map[int] bool

	//杠的栈记录
	Kong_stack bool

	//放杠人位置
	Kong_trigger_seat int

	//斗庄
	Dou_card int
	Dou_card_num int
	Is_dou_zhuang bool

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
	if len( this.Player_dict ) != this.Config.Max_chairs {
		return false
	}
	//有用户未准备
	for _,player := range this.Player_dict {
		if player.Status != define.EVENT_READY {
			return false
		}
	}
	this.Machine.Trigger(  )
	return false
}

func (self *Table) IsAllActed() bool {

	return true
}

//清除当前行为提示玩家记录
func (self *Table) ClearPrompts() {
	self.Player_prompts = make(map[int] bool)
}

//清除当前行为选择玩家记录
func (self *Table) ClearActions() {
	self.Player_actions = make(map[int] bool)
}

//过滤行为数组，去除低优先级操作
func (self *Table) FilterActions() {

}