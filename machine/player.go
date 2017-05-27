package machine

type Player struct {

	//用户id
	Uuid string

	//桌子
	Table Table

	//座位号
	Seat int

	//前一位置
	Prev_seat int

	//下一位置
	Next_seat int

	//当前事件
	Event string

	//状态机
	Machine *PlayerMachine

	//当前状态
	Status PlayerStatus

	//====================================总局数值====================================
	//总分
	Total int

	//杠动作分
	Kong_total int

	//明杠次数
	Kong_exposed_total int

	//暗杠次数
	Kong_concealed_total int

	//总获得分数
	Win_total_cnt int

	//自摸胡次数
	Win_draw_cnt int

	//点炮胡次数
	Win_discard_cnt int

	//点炮次数
	Pao_cnt int

	//是否房主
	Is_owner bool

	//====================================单轮数值====================================
	//获得分数
	Score int

	//杠动作分
	Kong_score int

	//手牌
	Cards_in_hand []int

	//碰杠牌
	Cards_group []int

	//出过的牌
	Cards_discard []int

	//可碰的牌
	Cards_pong []int

	//可明杠的牌
	Cards_kong_exposed []int

	//可暗杠的牌
	Cards_kong_concealed []int

	//可听的牌
	Cards_ready_hand []int

	//可胡的牌
	Cards_win []int

	//明杠次数
	Kong_exposed_cnt int

	//暗杠次数
	Kong_concealed_cnt int

	//过路杠次数
	Kong_pong_cnt int

	//放杠次数
	Kong_discard_cnt int

	//听牌提示
	Cards_dic []int

	//漏碰的牌
	Miss_pong_cards []int

	//漏胡的牌
	Miss_win_cards []int

	//过手胡分数
	Miss_win_card_score int

	//抓的牌
	Draw_card int

	//过路杠的牌
	Draw_kong_exposed_card int

	//胡的牌
	Win_card int

	//胡牌类型：点炮 自摸
	Win_type int

	//胡牌牌型
	Win_flag []string

	//提示
	Prompts []Prompt

	//动作
	Action Action

}

func (self *Player) Ready() {
	self.Machine.Trigger( &PlayerReadyStatus{} )
}