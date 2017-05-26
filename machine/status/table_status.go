package status

import "../../objects"

type TableReadyStatus struct{}
func (this *TableReadyStatus) Enter( table objects.Table ) {

	//定位上下家
	for seat,player := range table.Player_dict {
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
func (this *TableReadyStatus) NextStatus( table objects.Table ) {
	table.Machine.Trigger( DealState() )
}