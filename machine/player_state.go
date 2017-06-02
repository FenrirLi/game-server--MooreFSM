package machine

import (
	"log"
)

type PlayerState interface {
	Enter( player *Player )
	Execute( player *Player, event string )
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
func (this *PlayerInitState) Execute( player *Player, event string ) {
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
func (this *PlayerReadyState) Execute( player *Player, event string ) {
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
func (this *PlayerDealState) Execute( player *Player, event string ) {
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
	player.Table.Machine.CurrentState.Execute( player.Table, "TABLE_EVENT_PROMPT_DEAL" )
}


//===========================PlayerWaitState===========================
type PlayerWaitState struct {}
func (this *PlayerWaitState) Enter( player *Player ) {
	log_PlayerState( player, "enter", "wait" )
}
func (this *PlayerWaitState) Execute( player *Player, event string ) {
	log_PlayerState( player, "execute", "wait" )
	if interface_player_execute( player, event ) {
		switch event {

			case PlayerEvent["PLAYER_EVENT_OTHER_DISCARD"] :
				PlayerManagerCondition( player, "PLAYER_RULE_DISCARD" )

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