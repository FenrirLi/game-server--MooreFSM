package machine

type Action struct {

	//用户行为触发的状态
	State PlayerState

	//触发动作牌
	ActionCard int

	//手中参照牌
	ReferenceCard [] int

	//行为权重
	Weight int

}
