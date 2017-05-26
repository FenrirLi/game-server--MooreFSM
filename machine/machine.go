package machine

import (
	"./status"
)

type Machine struct {

	CurrentStatus status.Status

	LastStatus status.Status

}

func (self *Machine) Trigger( status status.Status ) {

}

func (self *Machine) BackToLastStatus() {

}

func (self *Machine) Execute() {

}

func (self *Machine) NextStatus() {

}
