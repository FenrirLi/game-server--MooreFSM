package handlers

import (
	"../teleport"
	"fmt"
	"../global"
	"strconv"
)

type EnterRoom struct {}

func (*EnterRoom) Process(receive *teleport.NetData) *teleport.NetData {

	fmt.Println("----------enter room-----------")
	//fmt.Println(receive.Body)
	//fmt.Println(receive.Operation)
	//fmt.Println(receive.From)
	//fmt.Println(receive.To)
	//fmt.Println(receive.Status)
	//fmt.Println(receive.Flag)

	table_id,_ := strconv.Atoi(receive.Body)
	fmt.Println(table_id)
	if table,ok := global.GLOBAL_TABLE[table_id]; ok {
		fmt.Println("enter room "+receive.Body+" success")
		fmt.Println(table.TableId)
	}else {

		return teleport.ReturnData("no such room","CreateRoomReturn",receive.From)
	}

	return nil
}