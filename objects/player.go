package objects

type Player struct {

	uuid string
	table Table
	seat int
	prev_seat int
	next_seat int

	is_online bool
	state = None
	self.vote_state = None
	self.vote_timer = None
	self.status = 0
self.event = None

self.machine = None
Machine(self)

# 累计数值
self.total = 0
self.kong_total = 0
self.kong_exposed_total = 0
self.kong_concealed_total = 0
self.win_total_cnt = 0
self.win_draw_cnt = 0
self.win_discard_cnt = 0
self.pao_cnt = 0
self.is_owner = 0

# 单局数值
self.score = 0
self.kong_score = 0
self.cards_in_hand = []
self.cards_group = []
self.cards_discard = []
self.cards_chow = []
self.cards_pong = []
self.cards_kong_exposed = []
self.cards_kong_concealed = []
self.cards_ready_hand = []
self.cards_draw_niao = []
self.cards_win = []
self.kong_exposed_cnt = 0
self.kong_concealed_cnt = 0
self.kong_discard_cnt = 0
self.kong_pong_cnt = 0
self.cards_dic = {}

# 漏碰的牌
self.miss_pong_cards = []
# 漏胡的牌
self.miss_win_cards = []
# 过手胡分数
self.miss_win_card_score = 0
self.draw_card = 0
self.draw_kong_exposed_card = 0
# 胡的牌
self.win_card = 0
# 胡牌类型：点炮 自摸
self.win_type = 0
# 胡牌牌型：将将胡 | 碰碰胡 | 七小对 | 。。。
self.win_flag = []
self.offline_cmd_stack = []

self.prompts = 0
# 提示ID
self.prompt_id = 0
# 动作
self.action_dict = {}
self.action_id = 0
self.action_weight = 0
self.action_rule = None
self.action_ref_cards = None
self.action_op_card = None

}