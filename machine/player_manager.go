package machine

import "log"

var PlayerRulesGroup = map[string][]PlayerRule{
	"PLAYER_RULE_READY": {},
	"PLAYER_RULE_DRAW": {},//[DrawConcealedKongRule(), DrawExposedKongRule(), DrawWinRule()],
	"PLAYER_RULE_DISCARD": {},//[DiscardExposedKongRule(), PongRule(), DiscardWinRule()],
	"PLAYER_RULE_DEAL": {},
	//PLAYER_RULE_CHOW: [],
	//PLAYER_RULE_PONG: [DrawConcealedKongRule(), DrawExposedKongRule()],
	"PLAYER_RULE_KONG": {}, //[QGWinRule()],
	//PLAYER_RULE_WIN: [],
	//PLAYER_RULE_NIAO: [],
}

func PlayerManagerCondition( player *Player, rule_group string ) {
	//依据检验的组对规则进行遍历
	if rules_array, ok := PlayerRulesGroup[rule_group]; ok {
		flag := false
		for _,rule := range rules_array {
			//满足规则则进行处理
			if rule.Condition( player ) {
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