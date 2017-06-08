package machine

type PlayerRule interface {
	//检验条件
	Condition( player *Player ) bool
	//进行处理
	Action( player *Player )
	//进行处理
	AddPrompt( player *Player, action Action )
}

func interface_player_rule_add_prompt( player *Player, action Action ){
	player.HasPrompt = true
	player.PromptId += 1
	player.ActionDict[ player.PromptId ] = action
}

//==================================玩家暗杠====================================
type PlayerConcealedKongRule struct {}
func (self *PlayerConcealedKongRule) Condition( player *Player ) bool {
	if player.Table.CardsRest.Len() <= 0 {
		return false
	} else {
		record := make(map[int]int)
		for _,v := range player.CardsInHand {
			if v == 0 {
				continue
			} else {
				record[v]++
			}
		}
		var action = Action{}
		for k,v := range record {
			if v == 4 {
				action.ActionId = PlayerAction["PLAYER_ACTION_KONG_CONCEALED"]
				action.ActionCard = k
				action.Weight = PlayerAction["PLAYER_ACTION_KONG_CONCEALED"]
				self.AddPrompt( player, action )
			}
		}
		return true
	}
}
func (self *PlayerConcealedKongRule) Action( player *Player ) {
	player.Machine.Trigger( &PlayerConcealedKongRuleState{} )
}
func (self *PlayerConcealedKongRule) AddPrompt( player *Player, action Action ) {
	interface_player_rule_add_prompt( player, action )
}