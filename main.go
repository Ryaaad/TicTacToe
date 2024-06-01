package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Cordination struct {
	i int
	j int
}

type model struct {
	board          [][]string
	cursor         Cordination
	turn           int
	SelectedOption string
	Menucusor      int
	Options        []string
	end            bool
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
			if m.SelectedOption == "" {
				if m.Menucusor > 0 {
					m.Menucusor--
				}
			} else {
				if m.cursor.i > 0 {
					m.cursor.i--
				}
			}

		case "down", "j":
			if m.SelectedOption == "" {
				if m.Menucusor < len(m.Options)-1 {
					m.Menucusor++
				}
			} else {
				if m.cursor.i < len(m.board)-1 {
					m.cursor.i++
				}
			}

		case "left":
			if m.SelectedOption != "" {
				if m.cursor.j > 0 {
					m.cursor.j--
				}
			}
		case "right":
			if m.SelectedOption != "" {
				if m.cursor.j < len(m.board[0])-1 {
					m.cursor.j++
				}
			}

		case "enter", " ":
			if m.SelectedOption == "" {
				m.SelectedOption = m.Options[m.Menucusor]
			} else {
				if m.SelectedOption == "Play vs Ai" {
					// m.UpdateAI()
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
				if m.SelectedOption == "Multiplayer" {
					// m.UpdateMultiplayer()
					if m.turn == 1 {
						if m.board[m.cursor.i][m.cursor.j] == "." {
							m.board[m.cursor.i][m.cursor.j] = "O"
							if Terminal(m.board) {
								m.end = true
							}
							m.turn = Turn(m.turn)
						}
					}
					if m.turn == 0 {
						if m.board[m.cursor.i][m.cursor.j] == "." {
							m.board[m.cursor.i][m.cursor.j] = "X"
							if Terminal(m.board) {
								m.end = true
							}
							m.turn = Turn(m.turn)
						}
					}
				}
			}

		}
	}
	return m, nil
}

func Style() []lipgloss.Style {
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("63")).
		Padding(1, 2).
		Margin(1, 0).
		Bold(true)
	return []lipgloss.Style{
		titleStyle,
	}
}

func (m model) View() string {
	if m.SelectedOption == "" {
		optionMenu := "\n"
		for index, option := range m.Options {
			cusor := " "
			if m.Menucusor == index {
				cusor = ">"
			}
			optionMenu += fmt.Sprintf("%s %s \n", cusor, option)
		}
		optionMenu += "\nPress q to quit.\n"
		return optionMenu
	} else {
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
		return s
	}
}

func initialModel() model {
	return model{
		board: [][]string{
			{".", ".", "."},
			{".", ".", "."},
			{".", ".", "."},
		},
		cursor:         struct{ i, j int }{i: 0, j: 0},
		turn:           1,
		SelectedOption: "",
		Options: []string{
			"Play vs Ai",
			"Multiplayer",
		},
		Menucusor: 0,
		end:       false,
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("there's been an error: %v", err)
		os.Exit(1)
	}
}
