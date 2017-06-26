package machine

import (
	"log"
	"reflect"
)

var PlayerRulesGroup = map[string][]PlayerRule{
	"PLAYER_RULE_READY": {},
	"PLAYER_RULE_DRAW": {&PlayerConcealedKongRule{},&PlayerPongKongRule{},&PlayerDrawWinRule{}},
	"PLAYER_RULE_DISCARD": {&PlayerPongRule{},&PlayerExposedKongRule{}},//[ DiscardWinRule()],
	"PLAYER_RULE_DEAL": {},
	"PLAYER_RULE_PONG":{&PlayerConcealedKongRule{},&PlayerPongKongRule{}},
	"PLAYER_RULE_KONG": {}, //[QGWinRule()],
	//PLAYER_RULE_WIN: [],
}

func PlayerManagerCondition( player *Player, rule_group string ) {
	//依据检验的组对规则进行遍历
	if rules_array, ok := PlayerRulesGroup[rule_group]; ok {
		flag := false
		for _,rule := range rules_array {
			//满足规则则进行处理
			if rule.Condition( player ) {
				log.Println("满足",reflect.TypeOf(rule).String())
				flag = true
			}
		}
		//规则检验有能够触发的，玩家进入提示状态
		if flag {
			player.Machine.Trigger( &PlayerPromptState{} )
		} else {
			log.Println("player rule manager next1")
			player.Machine.CurrentState.NextState( player )
		}
	} else {
		log.Println("Player Manager : rule_group Not Found")
	}

	return
}