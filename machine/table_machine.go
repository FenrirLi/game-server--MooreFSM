package machine

type TableMachine struct {

	Owner *Table

	CurrentState TableState

	LastState TableState

}

func NewTableMachine( table *Table, current TableState, last TableState ) TableMachine {
	return TableMachine{
		Owner: table,
		CurrentState: current,
		LastState: last,
	}
}

func (self *TableMachine) Trigger( state TableState ) {
	if self.CurrentState != nil {
		self.CurrentState.Exit( self.Owner )
		self.LastState = self.CurrentState
	}
	self.CurrentState = state
	self.CurrentState.Enter( self.Owner )
}

func (self *TableMachine) BackToLastState() {

}

func (self *TableMachine) Execute() {

}

func (self *TableMachine) NextState() {

}
