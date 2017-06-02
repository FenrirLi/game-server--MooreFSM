package machine

var PlayerEvent = map[string]string{
	//玩家准备事件
	"PLAYER_EVENT_READY":"PLAYER_EVENT_READY",
	//其他玩家出牌事件
	"PLAYER_EVENT_OTHER_DISCARD":"PLAYER_EVENT_OTHER_DISCARD",
	//其他玩家杠牌
	"PLAYER_EVENT_OTHER_KONG":"PLAYER_EVENT_OTHER_KONG",
}

var TableEvent = map[string]string{
	//通知坐庄
	"TABLE_EVENT_PROMPT_DEAL":"TABLE_EVENT_PROMPT_DEAL",
}