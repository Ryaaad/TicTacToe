package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Cordination struct {
	i int
	j int
}

type MinMaxResult struct {
	Value int
	Move  map[string]int
}

type model struct {
	board  [][]string
	cursor Cordination
	turn   int
	end    bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor.i > 0 {
				m.cursor.i--
			}

		case "down", "j":
			if m.cursor.i < len(m.board)-1 {
				m.cursor.i++
			}
		case "left":
			if m.cursor.j > 0 {
				m.cursor.j--
			}
		case "right":
			if m.cursor.j < len(m.board[0])-1 {
				m.cursor.j++
			}

		case "enter", " ":
			if m.turn == 1 {
				if m.board[m.cursor.i][m.cursor.j] == "." {
					m.board[m.cursor.i][m.cursor.j] = "O"
					if Terminal(m.board) {
						m.end = true
					}
					m.turn = Turn(m.turn)
					Move := MinMax(m.board, m.turn).Move
					m.board[Move["i"]][Move["j"]] = "X"
					if Terminal(m.board) {
						m.end = true
					}
					m.turn = Turn(m.turn)
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := ""
	if !m.end {
		// The header
		s = "What is you next move?\n\n"
		// Iterate over our choices
		for i := range m.board {
			for j := range m.board[0] {
				cursorleft := " "
				cursorright := " "
				if m.cursor.i == i && m.cursor.j == j {
					cursorleft = ">"
					cursorright = "<"
				}
				s += fmt.Sprintf("%s %s %s", cursorleft, m.board[i][j], cursorright)
			}
			s += "\n"
		}
	} else {
		s += "\n  Game Over ! \n"
		if WinX(m.board) {
			s += " \n X Won ! \n"
		}
		if WinO(m.board) {
			s += "\n O Won ! \n"
		}
		if !WinO(m.board) && !WinX(m.board) && Draw(m.board) {
			s += "\n Drawwww \n"
		}
	}
	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func initialModel() model {
	return model{
		board: [][]string{
			{".", ".", "."},
			{".", ".", "."},
			{".", ".", "."},
		},
		cursor: struct{ i, j int }{i: 0, j: 0},
		turn:   1,
		end:    false,
	}
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
func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
