package machine

type PlayerMachine struct {

	Owner *Player

	CurrentState PlayerState

	LastState PlayerState

}

func NewPlayerMachine( player *Player, current PlayerState, last PlayerState ) PlayerMachine {
	return PlayerMachine{
		Owner: player,
		CurrentState: current,
		LastState: last,
	}
}

func (self *PlayerMachine) Trigger( state PlayerState ) {
	if self.CurrentState != nil {
		self.CurrentState.Exit( self.Owner )
		self.LastState = self.CurrentState
	}
	self.CurrentState = state
	self.CurrentState.Enter( self.Owner )
}

func (self *PlayerMachine) BackToLastState() {

}

func (self *PlayerMachine) Execute() {

}

func (self *PlayerMachine) NextState() {

}
