package handlers

import (
	"../teleport"
	"fmt"
	"../machine"
	"../global"
	"strconv"
)

type CreateRoom struct {}

func (*CreateRoom) Process(receive *teleport.NetData) *teleport.NetData {

	fmt.Println("---------create room------------")
	fmt.Println(receive.Body)
	fmt.Println(receive.Operation)
	fmt.Println(receive.From)
	fmt.Println(receive.To)
	fmt.Println(receive.Status)
	fmt.Println(receive.Flag)

	table := machine.CreateTable( receive.From )
	new_machine := machine.NewTableMachine( table, nil, nil )
	table.Machine = &new_machine
	global.GLOBAL_TABLE[table.TableId] = &table
	fmt.Println(strconv.Itoa(table.TableId))

	return teleport.ReturnData(strconv.Itoa(table.TableId),"CreateRoomReturn")
}