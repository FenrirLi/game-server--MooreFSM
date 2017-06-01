package machine

import (
	"log"
	"reflect"
)

type PlayerStatus interface {
	Enter( player Player )
	Execute( player Player, event PlayerStatus )
	Exit( player Player )
}

//===========================PlayerInitStatus===========================
type PlayerInitStatus struct {}
func (this *PlayerInitStatus) Enter( player Player ) {}
func (this *PlayerInitStatus) Execute( player Player, event PlayerStatus ) {

	if event == PlayerEventStatus["EVENT_READY"] {

	} else {
		log.Println("init_status error call event "+reflect.TypeOf(event).Name())
	}

}
func (this *PlayerInitStatus) Exit( player Player ) {}


//===========================PlayerReadyStatus===========================
type PlayerReadyStatus struct {}
func (this *PlayerReadyStatus) Enter( player Player ) {
	//广播


	//检测桌子状态
	player.Table.IsAllReady()

}
func (this *PlayerReadyStatus) Execute( player Player, event PlayerStatus ) {}
func (this *PlayerReadyStatus) Exit( player Player ) {}