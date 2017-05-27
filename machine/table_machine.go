package machine

type TableMachine struct {

	Owner *Table

	CurrentStatus TableStatus

	LastStatus TableStatus

}

func NewTableMachine( table Table, current TableStatus, last TableStatus ) TableMachine {
	return TableMachine{
		Owner: &table,
		CurrentStatus: current,
		LastStatus: last,
	}
}

func (self *TableMachine) Trigger( status TableStatus ) {
	if self.CurrentStatus != nil {
		self.CurrentStatus.Exit( *self.Owner )
		self.LastStatus = self.CurrentStatus
	}
	self.CurrentStatus = status
	self.CurrentStatus.Enter( *self.Owner )
}

func (self *TableMachine) BackToLastStatus() {

}

func (self *TableMachine) Execute() {

}

func (self *TableMachine) NextStatus() {

}
