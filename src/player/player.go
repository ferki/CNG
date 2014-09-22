package player

const VERSION = "Default Go folding player"

func BetRequest(state *GameState) int {
    return 0
}

func Showdown(state *GameState) {
    
}

func IsPair(c1 *Card, c2 *Card) bool{
	return c1.Rank==c2.Rank;
}

func Raise(state *GameState, value float) int {
	return int(value*state.Pot);
}

func Version() string {
    return VERSION
}
