package objects

import (
	"time"
	"./config"
)

type Room struct {

	//房主
	owner string

	//创建时间
	create_time time.Time

	//房间配置
	config config.RoomConfig

}