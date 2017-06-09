package machine

import (
	"../proto"
	"../global"
)

//===========================玩家暗杠状态操作===========================
type PlayerConcealedKongRuleState struct {}
func (this *PlayerConcealedKongRuleState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "PlayerConcealedKongRuleState" )

	active_card := player.Action.ActionCard
	card_num := 4
	for _,v := range player.CardsInHand {
		if v == active_card {
			card_num --
		}
	}
	//有四张操作牌才能暗杠
	if card_num <= 0 {
		//手中移除4张牌
		num := 4
		for k,v := range player.CardsInHand {
			if v == active_card {
				player.CardsInHand[k] = 0
				num--
			}
			if num == 0 {
				break
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
		var request = &server_proto.ActionResponse{
			player.Uid,
			int32(active_card),
			ClientPlayerAction["KONG_CONCEALED"],
			[]int32{},
		}
		data := server_proto.MessageEncode( request )
		for _,player := range player.Table.PlayerDict{
			global.SERVER.Request(data, "ActionResponse", "action_response", player.Uid)
		}

		//去除出牌人记录
		player.Table.DiscardSeat = -1

		//自己进入下一状态
		player.Machine.NextState()
	}

}
func (this *PlayerConcealedKongRuleState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "PlayerConcealedKongRuleState" )
}
func (this *PlayerConcealedKongRuleState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "PlayerConcealedKongRuleState" )
}
func (this *PlayerConcealedKongRuleState) NextState( player *Player ) {
	log_PlayerState( player, "next", "PlayerConcealedKongRuleState" )
	//检测听牌
	//TODO

	player.Machine.Trigger( &PlayerDrawState{} )
}



//===========================玩家过路杠状态操作===========================
type PlayerPongKongRuleState struct {}
func (this *PlayerPongKongRuleState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "PlayerPongKongRuleState" )

	active_card := player.Action.ActionCard
	flag := 0
	for _,v := range player.CardsInHand {
		if v == active_card {
			flag++
			break
		}
	}
	for _,v := range player.CardsPong {
		if v == active_card {
			flag++
			break
		}
	}

	//手中有,已碰牌中也有才能过路杠
	if flag == 2 {
		for k,v := range player.CardsInHand {
			if v == active_card {
				player.CardsInHand[k] = 0
				break
			}
		}
		//明杠牌记录
		player.CardsKongExposed = append( player.CardsKongExposed, active_card )
		//玩家过路杠操作记录+1
		player.KongPongCnt++
		//算明杠分
		for _,p := range player.Table.PlayerDict{
			//杠牌人+3，其他人-1
			if p.Uid == player.Uid {
				p.KongScore += 3
			} else {
				p.KongScore -= 1
			}
		}

		//通知其他玩家杠牌操作
		var request = &server_proto.ActionResponse{
			player.Uid,
			int32(active_card),
			ClientPlayerAction["KONG_PONG"],
			[]int32{},
		}
		data := server_proto.MessageEncode( request )
		for _,player := range player.Table.PlayerDict{
			global.SERVER.Request(data, "ActionResponse", "action_response", player.Uid)
		}

		//去除出牌人记录
		player.Table.DiscardSeat = -1

		//如果可以点炮胡,检验抢杠胡等状态
		player.Table.ClearPrompts()
		player.Table.ClearActions()
		for _,p := range player.Table.PlayerDict{
			if p.Uid != player.Uid {
				p.Machine.CurrentState.Execute( p, PlayerEvent["PLAYER_EVENT_OTHER_KONG"], nil )
			}
		}
		//发现有用户有可进行的操作,当前用户切换到暂停状态
		if len( player.Table.PlayerPrompts ) > 0 {
			player.Machine.Trigger( &PlayerPauseState{} )
			return
		}

		//自己进入下一状态
		player.Machine.NextState()
	}

}
func (this *PlayerPongKongRuleState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "PlayerPongKongRuleState" )
}
func (this *PlayerPongKongRuleState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "PlayerPongKongRuleState" )
}
func (this *PlayerPongKongRuleState) NextState( player *Player ) {
	log_PlayerState( player, "next", "PlayerPongKongRuleState" )
	//检测听牌
	//TODO

	player.Machine.Trigger( &PlayerDrawState{} )
}


//===========================玩家明杠状态操作===========================
type PlayerExposedKongRuleState struct {}
func (this *PlayerExposedKongRuleState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "PlayerExposedKongRuleState" )

	active_card := player.Action.ActionCard
	card_num := 3
	for _,v := range player.CardsInHand {
		if v == active_card {
			card_num --
		}
	}
	//有三张操作牌才能明杠
	if card_num <= 0 {
		//手中移除3张牌
		num := 3
		for k,v := range player.CardsInHand {
			if v == active_card {
				player.CardsInHand[k] = 0
				num--
			}
			if num == 0 {
				break
			}
		}
		//玩家明杠操作记录+1，并加三分
		player.CardsKongExposed = append( player.CardsKongExposed, active_card )
		player.KongExposedCnt++
		player.KongScore += 3
		//放杠人放杠记录+1，并扣三分
		trigger_player := player.Table.PlayerDict[ player.Table.ActiveSeat ]
		trigger_player.KongScore -= 3
		trigger_player.KongDiscardCnt++

		//通知其他玩家杠牌操作
		var request = &server_proto.ActionResponse{
			player.Uid,
			int32(active_card),
			ClientPlayerAction["KONG_EXPOSED"],
			[]int32{},
		}
		data := server_proto.MessageEncode( request )
		for _,player := range player.Table.PlayerDict{
			global.SERVER.Request(data, "ActionResponse", "action_response", player.Uid)
		}

		//去除出牌人记录
		player.Table.DiscardSeat = -1

		//自己进入下一状态
		player.Machine.NextState()
	}

}
func (this *PlayerExposedKongRuleState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "PlayerExposedKongRuleState" )
}
func (this *PlayerExposedKongRuleState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "PlayerExposedKongRuleState" )
}
func (this *PlayerExposedKongRuleState) NextState( player *Player ) {
	log_PlayerState( player, "next", "PlayerExposedKongRuleState" )
	//检测听牌
	//TODO

	player.Machine.Trigger( &PlayerDrawState{} )
}



//===========================玩家碰状态操作===========================
type PlayerPongRuleState struct {}
func (this *PlayerPongRuleState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "PlayerPongRuleState" )

	//手中必需有两张相同牌
	active_card := player.Action.ActionCard
	card_num := 2
	for _,v := range player.CardsInHand {
		if v == active_card {
			card_num --
		}
	}
	if card_num <= 0 {
		//手牌去除两张
		card_num = 2
		for k,v := range player.CardsInHand {
			if v == active_card {
				card_num --
				player.CardsInHand[k] = 0
			}
			if card_num == 0 {
				break
			}
		}
		//记录碰牌
		player.CardsPong = append( player.CardsPong, active_card )

		//从出牌者手中去除出牌,并转换到等待状态
		discard_player := player.Table.PlayerDict[player.Table.ActiveSeat]
		e := discard_player.CardsDiscard.Front()
		discard_player.CardsDiscard.Remove(e)
		discard_player.Machine.Trigger( &PlayerWaitState{} )

		//去除出牌人记录
		player.Table.DiscardSeat = -1

		//切换当前操作人
		player.Table.ActiveSeat = player.Seat

		//听牌提示
		//TODO

		//通知所有玩家碰牌操作
		var request = &server_proto.ActionResponse{
			player.Uid,
			int32(active_card),
			ClientPlayerAction["PONG"],
			[]int32{},
		}
		data := server_proto.MessageEncode( request )
		for _,player := range player.Table.PlayerDict{
			global.SERVER.Request(data, "ActionResponse", "action_response", player.Uid)
		}


		//该状态下规则检测
		log_PlayerState( player, "PLAYER_RULE_PONG", "checking..." )
		PlayerManagerCondition( player, "PLAYER_RULE_PONG" )
	}


}
func (this *PlayerPongRuleState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "PlayerPongRuleState" )
}
func (this *PlayerPongRuleState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "PlayerPongRuleState" )
}
func (this *PlayerPongRuleState) NextState( player *Player ) {
	log_PlayerState( player, "next", "PlayerPongRuleState" )
	player.Machine.Trigger( &PlayerDiscardState{} )
}