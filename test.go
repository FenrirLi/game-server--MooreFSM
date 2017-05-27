package main

import (
	"./machine"
	"./global"
)

// 有标识符UID的demo，保证了客户端链接唯一性
func main() {
	table := machine.CreateTable( "1" )
	new_machine := machine.NewTableMachine( table, nil, nil )
	table.Machine = &new_machine
	global.GLOBAL_TABLE[111] = &table
}
