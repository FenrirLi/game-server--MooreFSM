package status

import (
	"../../objects"
	"../../define"
	"fmt"
)

//===========================PlayerInitStatus===========================
type PlayerInitStatus struct {}
func (this *PlayerInitStatus) Execute( player objects.Player, event string ) {

	if event == define.EVENT_READY {

	} else {
		fmt.Println("init_status error call event "+event)
	}

}


//===========================PlayerReadyStatus===========================
type PlayerReadyStatus struct {}
func (this *PlayerReadyStatus) enter( player objects.Player ) {
	//广播


	//检测桌子状态
	player.Table.IsAllReady()

}