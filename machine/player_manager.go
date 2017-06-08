package machine

import "log"

var PlayerRulesGroup = map[string][]PlayerRules{
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

//===========================PlayerRulesManager===========================
//type PlayerRulesManager struct {}
func PlayerManagerCondition( player *Player, rule_group string ) bool {
	//依据检验的组对规则进行遍历
	if rules_array, ok := PlayerRulesGroup[rule_group]; ok {
		for _,rule := range rules_array {
			//满足规则则进行处理
			if rule.Condition( *player ) {
				rule.Action( *player )
			}
		}
	} else {
		log.Println("Manager : rule_group Not Found")
	}

	player.Machine.CurrentState.NextState( player )
	return false
}