package machine

import (
	"log"
	"reflect"
	"container/list"
	"../proto"
	"../global"
)

type Player struct {

	//用户id
	Uid string

	//桌子
	Table *Table

	//座位号
	Seat int

	//前一位置
	PrevSeat int

	//下一位置
	NextSeat int

	//状态机
	Machine *PlayerMachine

	//是否在线
	OnlineState bool

	//====================================总局数值====================================
	//总分
	ScoreTotal int

	//杠动作分
	ScoreKongTotal int

	//明杠次数
	KongExposedTotal int

	//暗杠次数
	KongConcealedTotal int

	//暗杠次数
	KongPongTotal int

	//放杠次数
	KongDiscardTotal int

	//自摸胡次数
	WinDrawCnt int

	//点炮胡次数
	WinDiscardCnt int

	//点炮次数
	PaoCnt int

	//====================================单轮数值====================================
	//获得分数
	Score int

	//杠动作分
	KongScore int

	//手牌
	CardsInHand [14]int

	//碰杠牌
	CardsGroup []int

	//出过的牌
	CardsDiscard *list.List

	//可碰的牌
	CardsPong []int

	//可明杠的牌
	CardsKongExposed []int

	//可暗杠的牌
	CardsKongConcealed []int

	//可听的牌
	CardsReadyHand []int

	//可胡的牌
	CardsWin []int

	//明杠次数
	KongExposedCnt int

	//暗杠次数
	KongConcealedCnt int

	//过路杠次数
	KongPongCnt int

	//放杠次数
	KongDiscardCnt int

	//听牌提示
	CardsPrompt []int

	//漏碰的牌
	MissPongCards []int

	//漏胡的牌
	MissWinCards []int

	//过手胡分数
	MissWinCardScore int

	//抓的牌
	DrawCard int

	//过路杠的牌
	DrawKongExposedCard int

	//胡的牌
	WinCard int

	//胡牌类型：点炮 自摸
	WinType int

	//胡牌牌型
	WinFlag []string

	//是否有提示
	HasPrompt bool
	//提示自增id
	PromptId int

	//可执行动作集合
	ActionDict map[int]Action
	//选择的动作
	Action Action

}

func CreatePlayer( uid string, table *Table ) Player {
	var seat int
	//查找该用户是不是已经在成员列表
	flag := false
	for seat = 0; seat < table.Config.MaxChairs; seat++ {
		data, ok := table.PlayerDict[seat]
		if ok && data.Uid == uid {
			flag = true
			break
		}
	}
	//不在成员列表则取最低座位号
	if !flag {
		for seat = 0; seat < table.Config.MaxChairs; seat++ {
			_, ok := table.PlayerDict[seat]
			if !ok {
				break
			}
		}
	}

	return Player{
		Uid: uid,
		Table: table,
		//需要修正 @debug
		Seat: seat,
		CardsDiscard:list.New(),
	}
}

func (self *Player) Ready() {
	self.Machine.Trigger( &PlayerReadyState{} )
}

//去除玩家提示信息
func ( self *Player ) DelPrompt() {
	self.HasPrompt = false
	self.PromptId = 0
	self.ActionDict = map[int]Action{}
}

//去除玩家动作信息
func ( self *Player ) DelAction() {
	self.Action = Action{}
}

//玩家单轮数据重置
func ( self *Player ) initRound() {
	//获得分数
	self.Score = 0
	//杠动作分
	self.KongScore = 0
	//手牌
	self.CardsInHand = [14]int{}
	//碰杠牌
	self.CardsGroup = []int{}
	//出过的牌
	self.CardsDiscard = list.New()
	//可碰的牌
	self.CardsPong = []int{}
	//可明杠的牌
	self.CardsKongExposed = []int{}
	//可暗杠的牌
	self.CardsKongConcealed = []int{}
	//可听的牌
	self.CardsReadyHand = []int{}
	//可胡的牌
	self.CardsWin = []int{}
	//明杠次数
	self.KongExposedCnt = 0
	//暗杠次数
	self.KongConcealedCnt = 0
	//过路杠次数
	self.KongPongCnt = 0
	//放杠次数
	self.KongDiscardCnt = 0
	//听牌提示
	self.CardsPrompt = []int{}
	//漏碰的牌
	self.MissPongCards = []int{}
	//漏胡的牌
	self.MissWinCards = []int{}
	//过手胡分数
	self.MissWinCardScore = 0
	//抓的牌
	self.DrawCard = 0
	//过路杠的牌
	self.DrawKongExposedCard = 0
	//胡的牌
	self.WinCard = 0
	//胡牌类型：点炮 自摸
	self.WinType = 0
	//胡牌牌型
	self.WinFlag = []string{}
	//是否有提示
	self.HasPrompt = false
	//提示自增id
	self.PromptId = 0
	//可执行动作集合
	self.ActionDict = map[int]Action{}
	//选择的动作
	self.Action = Action{}
}

//出牌
func ( self *Player ) Discard( card int ) {
	log.Println("      ",self.Uid,"出牌",card)
	//如果用户是在没有处理操作提示的情况下出的牌
	if reflect.DeepEqual( self.Machine.CurrentState, &PlayerPromptState{} ) {
		log.Println("没有处理提示，直接出牌")
		//清除桌子提示
		self.Table.ClearPrompts()
		//立刻处理掉玩家提示状态下的状态切换
		self.Machine.NextState()
	}

	//出牌
	self.Table.DiscardSeat = self.Seat
	flag := false
	for k,v := range self.CardsInHand {
		if v == card {
			self.CardsInHand[k] = 0
			flag = true
			break
		}
	}
	//不存在的手牌
	if( flag == false ) {
		log.Println("      ERROR:",self.Uid,"没有这张牌",card)
		return
	}
	log.Println(self.CardsInHand)
	self.CardsDiscard.PushFront(card)
	self.Table.ActiveCard = card

	//清除过手胡牌记录
	self.MissPongCards = []int{}
	self.MissWinCards = []int{}
	self.MissWinCardScore = 0

	//给所有人发出牌通知
	var request = &server_proto.ActionResponse{
		self.Uid,
		int32(card),
		ClientPlayerAction["DISCARD"],
		[]int32{},
	}
	data := server_proto.MessageEncode( request )
	for _,player := range self.Table.PlayerDict{
		global.SERVER.Request(data, "ActionResponse", "action_response", player.Uid)
	}

	//用户检测听牌状态
	//TODO

	//其他玩家执行"他人出牌"判定
	for k,player := range self.Table.PlayerDict{
		if player.Uid == self.Uid {
			continue
		}
		player.Machine.CurrentState.Execute( self.Table.PlayerDict[k], PlayerEvent["PLAYER_EVENT_OTHER_DISCARD"], nil )
	}

	//杠上操作记录清空
	self.Table.KongStack = false
	self.Table.KongTriggerSeat = -1

	//他人执行判定后，可能出现操作，因此进入等待状态
	if len(self.Table.PlayerPrompts) > 0 {
		self.Machine.Trigger( &PlayerPauseState{} )
	} else {

		log.Println("self machine next")
		self.Machine.NextState()
	}

}

//玩家操作
func ( self *Player ) ActionSelect( select_id int ) {
	//选了过
	if select_id == 0 {
		log.Println("    ----",self.Uid," PASSED PROMPT----")
	} else {
		if action,ok := self.ActionDict[select_id]; ok {
			self.Action = action
		} else {
			log.Println("ERROR: ----NO SUCH PROMPT----")
		}
	}
	self.Table.PlayerActions = append( self.Table.PlayerActions, self.Seat )

	//过掉低优先级操作
	//TODO
	//for _,seat := range self.Table.PlayerPrompts {
	//	//不和自己比较
	//	if seat == self.Seat {
	//		continue
	//	}
	//	//已经操作完毕的用户跳过
	//	if player,ok := self.Table.PlayerDict[seat]; ok {
	//		if player.Action.Weight > 0 {
	//			continue
	//		}
	//		//低优先级的未操作用户过滤掉
	//		max_weight := 0
	//		for _,action := range player.ActionDict {
	//			if action.Weight > max_weight{
	//				max_weight = action.Weight
	//			}
	//		}
	//		if max_weight < self.Action.Weight {
	//
	//		}
	//	}
	//}

	//相同优先级处理
	//TODO

	//桌子判定是否所有人已选择操作完毕
	self.Table.CheckAllActed()

}