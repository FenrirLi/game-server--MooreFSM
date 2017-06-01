package machine

type PlayerMachine struct {

	Owner *Player

	CurrentStatus PlayerStatus

	LastStatus PlayerStatus

}

func NewPlayerMachine( player Player, current PlayerStatus, last PlayerStatus ) PlayerMachine {
	return PlayerMachine{
		Owner: &player,
		CurrentStatus: current,
		LastStatus: last,
	}
}

func (self *PlayerMachine) Trigger( status PlayerStatus ) {
	if self.CurrentStatus != nil {
		self.CurrentStatus.Exit( *self.Owner )
		self.LastStatus = self.CurrentStatus
	}
	self.CurrentStatus = status
	self.CurrentStatus.Enter( *self.Owner )
}

func (self *PlayerMachine) BackToLastStatus() {

}

func (self *PlayerMachine) Execute() {

}

func (self *PlayerMachine) NextStatus() {

}
