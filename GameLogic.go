package main

func Terminal(S [][]string) bool {
	if WinX(S) {
		return true
	}
	if WinO(S) {
		return true
	}
	if Draw(S) {
		return true
	}
	return false
}

func Turn(turn int) int {
	return 1 - turn
}

func WinX(S [][]string) bool {
	for i := 0; i < len(S); i++ {
		if S[i][0] == S[i][1] && S[i][1] == S[i][2] && S[i][0] == "X" {
			return true
		}
	}
	for j := 0; j < len(S); j++ {
		if S[0][j] == S[1][j] && S[1][j] == S[2][j] && S[0][j] == "X" {
			return true
		}
	}
	if S[0][0] == S[1][1] && S[1][1] == S[2][2] && S[0][0] == "X" {
		return true
	}
	if S[0][2] == S[1][1] && S[1][1] == S[2][0] && S[0][2] == "X" {
		return true
	}

	return false
}

func WinO(S [][]string) bool {
	for i := 0; i < len(S); i++ {
		if S[i][0] == S[i][1] && S[i][1] == S[i][2] && S[i][0] == "O" {
			return true
		}
	}
	for j := 0; j < len(S); j++ {
		if S[0][j] == S[1][j] && S[1][j] == S[2][j] && S[0][j] == "O" {
			return true
		}
	}
	if S[0][0] == S[1][1] && S[1][1] == S[2][2] && S[0][0] == "O" {
		return true
	}
	if S[0][2] == S[1][1] && S[1][1] == S[2][0] && S[0][2] == "O" {
		return true
	}

	return false
}

func Draw(S [][]string) bool {
	for i := 0; i < len(S); i++ {
		for j := 0; j < len(S[0]); j++ {
			if S[i][j] == "." {
				return false
			}
		}
	}
	return true
}
