package machine

func ReadyHand( hand_cards [14]int ) []int{

	index := -1
	for k,v := range hand_cards{
		if v == 0 {
			index = k
		}
	}

	result := []int{}

	if index == -1 {
		return result
	}

	for k,_ := range Cards{
		hand_cards[index] = k
		if WinCheck(hand_cards) {
			result = append(result, k)
		}
	}

	return result
}

func WinCheck( hand_cards [14]int ) bool {

	cards := make(map[int]int)

	for _,v := range hand_cards{
		cards[v]++
	}

	if TryWin( false, cards ) {
		return true
	} else {
		return false
	}

}

func CountCard( self_cards map[int]int , card int ) int {
	if v, ok := self_cards[card]; ok {
		return v
	} else {
		return 0
	}
}

func TryWin( self_has_pair bool, self_cards map[int]int ) bool {

	rem_pair := self_has_pair
	rem_cards := make(map[int]int)
	for k,v := range self_cards{
		rem_cards[k] = v
	}


	//取最小牌
	active_card := 99
	for k,v := range self_cards{
		if k < active_card && v > 0 {
			active_card = k
		}
	}

	if active_card == 99 {
		if self_has_pair {
			return true
		} else {
			return false
		}
	}

	if TryPair( &self_has_pair, &self_cards, active_card ) {
		if TryWin( self_has_pair,self_cards ) {
			return true
		}
	}
	self_has_pair = rem_pair
	self_cards = rem_cards

	if TryTriplets( &self_cards, active_card ) {
		if TryWin( self_has_pair,self_cards ) {
			return true
		}
	}

	//if self.TrySequence( active_card ) {
	//	if self.TryWin() {
	//		return true
	//	}
	//}

	return false
}

func TryPair( self_has_pair *bool, self_cards *(map[int]int), card int ) bool {
	if *self_has_pair {
		return false
	}
	if CountCard( *self_cards, card ) >= 2 {
		(*self_cards)[card] -= 2
		*self_has_pair = true
		return true
	} else {
		return false
	}

}

func TryTriplets( self_cards *(map[int]int), card int ) bool {
	if CountCard( *self_cards , card ) >= 3 {
		(*self_cards)[card] -= 3
		return true
	} else {
		return false
	}
}

//func TrySequence( card int ) bool {
//
//}
