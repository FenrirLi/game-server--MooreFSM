package objects

import (
	"time"
	"./config"
	"container/list"
)


type TableInfo struct {

	//桌子id
	table_id string

	//创建人
	owner string

	//创建时间
	create_time time.Time

	//桌子配置
	config config.TableConfig

	//玩家
	player_dict map[int] Player

	//状态机
	machine string
	state string

	//庄家位置
	dealer_seat int

	//当前行动人位置
	active_seat int

	//当前出牌
	active_card int

	//当前出牌位置
	discard_seat int

	//event map[int]

	//剩余卡牌
	cards_rest list.List

	//获得提示的玩家位置
	player_prompts map[int] bool

	//做出行为选择的玩家位置
	player_actions map[int] bool

	//牌局记录
	replay list.List

	//胡牌类型：流局，自摸胡，点炮胡，一炮多响
	win_type int

	//当前局数
	cur_round int

	//赢家
	winner_list map[int] bool

	//输家
	loser_list map[int] bool

	//杠的栈记录
	kong_stack bool

	//放杠人位置
	kong_trigger_seat int

	//斗庄
	dou_card int
	dou_card_num int
	is_dou_zhuang bool

}

type Table interface {
	//初始化桌面数据
	InitTable() Table

	//根据当前桌面数据生成下局桌面数据
	NextRound() Table

	//玩家参与
	Enter()

	//桌子解散
	DisMiss()

	//是否都做了行动
	IsAllActed() bool

	//清除当前有提示的玩家记录
	ClearPrompts()

	//清除当前选择了行动的玩家记录
	ClearActions()

	//过滤低优先级行为
	FilterActions()

}

func new() Table {
	return &TableInfo{

	}
}

func (self *TableInfo) Enter() {
	//TODO

}

func (self *TableInfo) DisMiss() {
	//TODO

}

func (self *TableInfo) InitTable() Table {
	//TODO
	return self
}

func (self *TableInfo) NextRound() Table {
	//TODO
	return self
}

func (self *TableInfo) IsAllActed() bool {

	return true
}

//清除当前行为提示玩家记录
func (self *TableInfo) ClearPrompts() {
	self.player_prompts = make(map[int] bool)
}

//清除当前行为选择玩家记录
func (self *TableInfo) ClearActions() {
	self.player_actions = make(map[int] bool)
}

//过滤行为数组，去除低优先级操作
func (self *TableInfo) FilterActions() {

}