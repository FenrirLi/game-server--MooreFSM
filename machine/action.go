package machine

type Action struct {

	//用户触发的事件
	ActionId int

	//触发动作牌
	ActionCard int

	//手中参照牌
	ReferenceCard []int32

	//行为权重
	Weight int

}
