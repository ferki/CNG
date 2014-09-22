package player

const VERSION = "v1.0.0"

func BetRequest(state *GameState) int {
	if IsPair(PLAYER) {
		return Raise(0.67);
	}
    else {
    	return 0;
    }
}

func Showdown(state *GameState) {
    
}

func IsPair_helper(c1 *Card, c2 *Card) bool {
	return (c1.Rank==c2.Rank);
}

func IsPair(p Player) bool {
	return (IsPair_helper(p.HoleCards[0], p.HoleCards[1]));
}

func Raise(state *GameState, value float) int {
	// return int(value*state.Pot);
	return (ceil((value*pot)/10))*10;
}

func HasAce(p Player) bool{
	for i=0; i<len(p.HoleCards); i++
		{
			if Contains(p.HoleCards.Rank, "A") {
				return true;
			}
		}
	return false;
}

func Version() string {
    return VERSION;
}