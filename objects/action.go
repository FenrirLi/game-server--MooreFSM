package objects

type Action struct {

	//行为id
	Id int

	//权重
	Weight int

	//触发动作牌
	Action_card int

	//规则
	Rule string

	//手中参照牌
	Reference_card [] int
}
