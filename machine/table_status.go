package machine

type TableStatus interface {
	Enter( table Table )
	Execute( table Table, event TableStatus )
	Exit( table Table )
	NextStatus( table Table )
}


//===========================TableReadyStatus===========================
type TableReadyStatus struct{}
func (this *TableReadyStatus) Enter( table Table ) {

	//定位上下家
	for seat,player := range table.PlayerDict {
		//下家
		next_seat := seat + 1;
		if  next_seat >= table.Config.Max_chairs {
			next_seat -= table.Config.Max_chairs
		}
		player.Next_seat = next_seat

		//上家
		prev_seat := seat - 1
		if  prev_seat < 0 {
			prev_seat += table.Config.Max_chairs
		}
		player.Prev_seat = prev_seat
	}

}
func (this *TableReadyStatus) Execute( table Table, event TableStatus ) {}
func (this *TableReadyStatus) Exit( table Table ) {}
func (this *TableReadyStatus) NextStatus( table Table ) {
	//table.Machine.Trigger( DealState() )
}