package machine

//===========================玩家暗杠状态操作===========================
type PlayerConcealedKongRuleState struct {}
func (this *PlayerConcealedKongRuleState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "ConcealedKongRule" )

	active_card := player.Action.ActionCard
	card_num := 4
	for _,v := range player.CardsInHand {
		if v == active_card {
			card_num --
		}
	}
	//有四张操作牌才能暗杠
	if card_num == 0 {
		for k,v := range player.CardsInHand {
			if v == active_card {
				player.CardsInHand[k] = 0
			}
		}
		//暗杠牌记录
		player.CardsKongConcealed = append( player.CardsKongConcealed, active_card )
		//玩家暗杠操作记录+1
		player.KongConcealedCnt ++
		//算暗杠分
		for _,p := range player.Table.PlayerDict{
			//杠牌人+6，其他人-2
			if p.Uid == player.Uid {
				p.KongScore += 6
			} else {
				p.KongScore -= 2
			}
		}

		//通知其他玩家杠牌操作
		//TODO

		//此轮没有人出牌
		player.Table.DiscardSeat = -1

		//自己进入下一状态
		player.Machine.NextState()
	}

}
func (this *PlayerConcealedKongRuleState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "ConcealedKongRule" )
}
func (this *PlayerConcealedKongRuleState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "ConcealedKongRule" )
}
func (this *PlayerConcealedKongRuleState) NextState( player *Player ) {
	log_PlayerState( player, "next", "ConcealedKongRule" )
	//检测听牌
	//TODO

	player.Machine.Trigger( &PlayerDrawState{} )
}
