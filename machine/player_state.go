package machine

import (
	"log"
	"../proto"
	"../global"
	"reflect"
)

type PlayerState interface {
	Enter( player *Player )
	Execute( player *Player, event string, request_body []byte )
	Exit( player *Player )
	NextState( player *Player )
}

func log_PlayerState( player *Player, act string, state string ) {
	log.Printf("    ----player %s %s %s ----",player.Uid,act,state)
}

func interface_player_execute( player *Player, event string ) bool {
	//当前状态下执行event事件
	define_event, ok := PlayerEvent[event]
	if ok && event == define_event {
		log.Printf("      player --%s-- ACTIVE -- %s ",player.Uid,event)
	} else {
		log.Println("PlayerInitState error call event "+event)
	}
	return ok
}

//===========================PlayerInitState===========================
type PlayerInitState struct {}
func (this *PlayerInitState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "init" )
}
func (this *PlayerInitState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "init" )
	interface_player_execute( player, event )
}
func (this *PlayerInitState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "init" )
}
func (this *PlayerInitState) NextState( player *Player ) {
	log_PlayerState( player, "next", "init" )
}


//===========================PlayerReadyState===========================
type PlayerReadyState struct {}
func (this *PlayerReadyState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "ready" )
	//广播


	//检测桌子状态
	player.Table.IsAllReady()

}
func (this *PlayerReadyState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "ready" )
	interface_player_execute( player, event )
}
func (this *PlayerReadyState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "ready" )
}
func (this *PlayerReadyState) NextState( player *Player ) {
	log_PlayerState( player, "next", "ready" )
}


//===========================PlayerDealState===========================
type PlayerDealState struct {}
func (this *PlayerDealState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "deal" )
	//该状态下规则检测
	log_PlayerState( player, "PLAYER_RULE_DEAL", "checking..." )
	PlayerManagerCondition( player, "PLAYER_RULE_DEAL" )
}
func (this *PlayerDealState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "deal" )
	interface_player_execute( player, event )
}
func (this *PlayerDealState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "deal" )
}
func (this *PlayerDealState) NextState( player *Player ) {
	log_PlayerState( player, "next", "deal" )
	//玩家进入等待状态
	player.Machine.Trigger( &PlayerWaitState{} )
	//桌子触发事件
	player.Table.Machine.CurrentState.Execute( player.Table, "TABLE_EVENT_PROMPT_DEAL", nil )
}


//===========================PlayerWaitState===========================
type PlayerWaitState struct {}
func (this *PlayerWaitState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "wait" )
}
func (this *PlayerWaitState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "wait" )
	if interface_player_execute( player, event ) {
		switch event {

			//他人出牌
			case PlayerEvent["PLAYER_EVENT_OTHER_DISCARD"] :
				PlayerManagerCondition( player, "PLAYER_RULE_DISCARD" )
			//他人杠牌
			case PlayerEvent["PLAYER_EVENT_OTHER_KONG"] :
				PlayerManagerCondition( player, "PLAYER_RULE_KONG" )

			default:
				log.Println("----- no such event ",event," ----- ")
		}
	}
}
func (this *PlayerWaitState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "wait" )
}
func (this *PlayerWaitState) NextState( player *Player ) {
	log_PlayerState( player, "next", "wait" )
}


//===========================PlayerDrawState===========================
type PlayerDrawState struct {}
func (this *PlayerDrawState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "draw" )

	//清除玩家上一个动作数据
	player.MissPongCards = []int{}
	player.MissWinCards = []int{}
	player.MissWinCardScore = 0

	//清除桌子提示和行为数据
	player.Table.ClearPrompts()
	player.Table.ClearActions()

	//当前用户听牌计算
	//TODO

	//用户抓牌
	e := player.Table.CardsRest.Front()
	player.Table.CardsRest.Remove(e)
	draw_card := e.Value.(int)
	player.DrawCard = draw_card
	for k,v := range player.CardsInHand {
		if v == 0 {
			player.CardsInHand[k] = draw_card
			break
		}
	}

	//推送抓牌消息
	var request = &server_proto.DrawCardResponse{}
	var data []byte
	request.Card = int32(draw_card)
	data = server_proto.MessageEncode( request )
	global.SERVER.Request(data, "DrawResponse", "draw_response", player.Uid)

	//该状态下规则检测
	log_PlayerState( player, "PLAYER_RULE_DRAW", "checking..." )
	PlayerManagerCondition( player, "PLAYER_RULE_DRAW" )
}
func (this *PlayerDrawState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "draw" )
	interface_player_execute( player, event )
}
func (this *PlayerDrawState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "draw" )
}
func (this *PlayerDrawState) NextState( player *Player ) {
	log_PlayerState( player, "next", "draw" )
	player.Machine.Trigger( &PlayerDiscardState{} )

	log.Println("back to draw")
}


//===========================PlayerDiscardState===========================
type PlayerDiscardState struct {}
func (this *PlayerDiscardState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "discard" )

	//已经进入出牌步骤，清除之前操作记录
	player.Table.ClearPrompts()
	player.Table.ClearActions()


	log.Println("discard end")
}
func (this *PlayerDiscardState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "discard" )
	if interface_player_execute( player, event ) {
		switch event {

			//出牌
			case PlayerEvent["PLAYER_EVENT_DISCARD"] :
				// 进行解码
				request := &server_proto.DiscardRequest{}
				server_proto.MessageDecode( request_body, request )

				player.Discard( int(request.Card) )

			default:
				log.Println("----- no such event ",event," ----- ")
		}
	}
}
func (this *PlayerDiscardState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "discard" )
}
func (this *PlayerDiscardState) NextState( player *Player ) {
	log_PlayerState( player, "next", "discard" )
	player.Machine.Trigger( &PlayerWaitState{} )
	player.Table.Machine.CurrentState.Execute( player.Table, TableEvent["TABLE_EVENT_STEP"], nil )
}

//===========================PlayerPromptState===========================
type PlayerPromptState struct {}
func (this *PlayerPromptState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "prompt" )
	//往桌子添加玩家提示记录
	player.Table.PlayerPrompts = append( player.Table.PlayerPrompts, player.Seat )

	//给玩家发送提示
	request := &server_proto.ActionPrompt{}
	action_proto := &server_proto.Action{}
	for k,v := range player.ActionDict {
		action_proto.SelectId = int32(k)
		action_proto.ActionId = int32(v.ActionId)
		action_proto.ActionCard = int32(v.ActionCard)
		action_proto.RefCards = v.ReferenceCard
		action_proto.Weight = int32(v.Weight)
		request.Action = append(request.Action, action_proto)
	}
	data := server_proto.MessageEncode( request )
	global.SERVER.Request( data, "ActionPrompt","action_prompt_flag", player.Uid )


}
func (this *PlayerPromptState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "prompt" )
	if interface_player_execute( player, event ) {
		switch event {

		//操作
		case PlayerEvent["PLAYER_EVENT_ACTION"] :
			// 进行解码
			request := &server_proto.ActionSelectRequest{}
			server_proto.MessageDecode( request_body, request )

			player.ActionSelect( int(request.SelectId) )

		//出牌
		case PlayerEvent["PLAYER_EVENT_DISCARD"] :
			// 进行解码
			request := &server_proto.DiscardRequest{}
			server_proto.MessageDecode( request_body, request )

			player.Discard( int(request.Card) )

		default:
			log.Println("----- no such event ",event," ----- ")
		}
	}
}
func (this *PlayerPromptState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "prompt" )
}
func (this *PlayerPromptState) NextState( player *Player ) {
	log_PlayerState( player, "next", "prompt" )
	if reflect.DeepEqual( player.Machine.CurrentState, &PlayerWaitState{} ) {
		player.Machine.BackToLastState()
	} else {

		log.Println("player prompt execute next")
		player.Machine.LastState.NextState(player)
	}
}


//===========================PlayerPauseState===========================
type PlayerPauseState struct {}
func (this *PlayerPauseState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "pause" )
}
func (this *PlayerPauseState) Execute( player *Player, event string, request_body []byte ) {
	log_PlayerState( player, "execute", "pause" )
	if interface_player_execute( player, event ) {

	}
}
func (this *PlayerPauseState) Exit( player *Player ) {
	log_PlayerState( player, "exit", "pause" )
}
func (this *PlayerPauseState) NextState( player *Player ) {
	log_PlayerState( player, "next", "pause" )
	player.Machine.LastState.NextState( player )
}