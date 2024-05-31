package main

type MinMaxResult struct {
	Value int
	Move  map[string]int
}

func Evaluate(S [][]string) int {
	if WinX(S) {
		return 1
	}
	if WinO(S) {
		return -1
	}
	return 0
}

func MinMax(S [][]string, turn int) MinMaxResult {
	if Terminal(S) {
		return MinMaxResult{Value: Evaluate(S), Move: nil}
	}

	Actions := Actions(S, turn)
	var bestMove map[string]int

	if turn == 0 {
		value := -2
		for i := 0; i < len(Actions); i++ {
			result := MinMax(Result(Actions[i], S, turn), 1)
			if result.Value > value {
				value = result.Value
				bestMove = Actions[i]
			}
		}
		return MinMaxResult{Value: value, Move: bestMove}
	} else {
		value := 2
		for i := 0; i < len(Actions); i++ {
			result := MinMax(Result(Actions[i], S, turn), 0)
			if result.Value < value {
				value = result.Value
				bestMove = Actions[i]
			}
		}
		return MinMaxResult{Value: value, Move: bestMove}
	}
}

func Result(Action map[string]int, S [][]string, turn int) [][]string {
	Scopy := make([][]string, len(S))
	for i := range S {
		Scopy[i] = make([]string, len(S[i]))
		copy(Scopy[i], S[i])
	}
	if turn == 0 {
		Scopy[Action["i"]][Action["j"]] = "X"
	}
	if turn == 1 {
		Scopy[Action["i"]][Action["j"]] = "O"
	}
	return Scopy
}

func Actions(S [][]string, turn int) []map[string]int {
	var Possible_Actions []map[string]int
	for i := 0; i < len(S); i++ {
		for j := 0; j < len(S[0]); j++ {
			if S[i][j] == "." {
				action := map[string]int{"i": i, "j": j}
				Possible_Actions = append(Possible_Actions, action)
			}
		}
	}
	return Possible_Actions
}
