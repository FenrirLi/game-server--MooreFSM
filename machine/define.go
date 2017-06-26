package machine


var GLOBAL_TABLE = make(map[int]*Table)

var GLOBAL_USER = make(map[string]*Player)

var PlayerEvent = map[string]string{
	//玩家准备事件
	"PLAYER_EVENT_READY":"PLAYER_EVENT_READY",
	//其他玩家出牌事件
	"PLAYER_EVENT_OTHER_DISCARD":"PLAYER_EVENT_OTHER_DISCARD",
	//其他玩家杠牌
	"PLAYER_EVENT_OTHER_KONG":"PLAYER_EVENT_OTHER_KONG",
	//出牌
	"PLAYER_EVENT_DISCARD":"PLAYER_EVENT_DISCARD",
	//操作
	"PLAYER_EVENT_ACTION":"PLAYER_EVENT_ACTION",
}

var PlayerAction = map[string]int{
	//碰牌
	"PLAYER_ACTION_PONG":1,
	//暗杠
	"PLAYER_ACTION_KONG_CONCEALED":2,
	//明杠
	"PLAYER_ACTION_KONG_EXPOSED":3,
	//过路杠
	"PLAYER_ACTION_KONG_PONG":4,
	//自摸
	"PLAYER_ACTION_WIN_DRAW":5,
	//点炮胡
	"PLAYER_ACTION_WIN_DISCARD":6,
}

var ClientPlayerAction = map[string]string {
	//抓牌
	"DRAW":"DRAW",
	//出牌
	"DISCARD":"DISCARD",
	//碰牌
	"PONG":"PONG",
	//明杠牌
	"KONG_EXPOSED":"KONG_EXPOSED",
	//暗杠牌
	"KONG_CONCEALED":"KONG_CONCEALED",
	//过路杠牌
	"KONG_PONG":"KONG_PONG",
	//自摸胡
	"WIN_DRAW":"WIN_DRAW",
	//点炮胡
	"WIN_DISCARD":"WIN_DISCARD",
}

var PlayerActionRule = map[int]PlayerRule{
	//碰牌
	1:&PlayerPongRule{},
	//暗杠
	2:&PlayerConcealedKongRule{},
	//明杠
	3:&PlayerExposedKongRule{},
	//过路杠
	4:&PlayerPongKongRule{},
	//自摸
	5:&PlayerDrawWinRule{},
	//点炮胡
	6:&PlayerDiscardWinRule{},
}

var TableEvent = map[string]string{
	//等待玩家准备
	"TABLE_EVENT_WAIT_READY":"TABLE_EVENT_WAIT_READY",
	//通知坐庄
	"TABLE_EVENT_PROMPT_DEAL":"TABLE_EVENT_PROMPT_DEAL",
	//步骤
	"TABLE_EVENT_STEP":"TABLE_EVENT_STEP",
	//结束
	"TABLE_EVENT_END":"TABLE_EVENT_END",
}

var Cards = map[int]string{
	1:"东风",
	2:"西风",
	3:"南风",
	4:"北风",
	5:"红中",
	6:"发财",
	7:"白板",

	11:"1条",
	12:"2条",
	13:"3条",
	14:"4条",
	15:"5条",
	16:"6条",
	17:"7条",
	18:"8条",
	19:"9条",

	21:"1筒",
	22:"2筒",
	23:"3筒",
	24:"4筒",
	25:"5筒",
	26:"6筒",
	27:"7筒",
	28:"8筒",
	29:"9筒",

	31:"1万",
	32:"2万",
	33:"3万",
	34:"4万",
	35:"5万",
	36:"6万",
	37:"7万",
	38:"8万",
	39:"9万",
}

var WinTypes = map[string]int{
	"WIN_DISCARD_ONE":1,
	"WIN_DISCARD_MORE":2,
	"WIN_DRAW":3,
}

var Scores = map[string]int{
	"QI_XIAO_DUI":1,
	"QING_YI_SE":1,
	"PONG_PONG_HU":1,
}