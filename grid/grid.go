package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	GRID_SIZE = 10
)

type grid struct {
	x      int
	y      int
	symbol string
}

func initGrid() grid {
	return grid{
		x: 0, y: 0, symbol: ">",
	}
}

func (m grid) Init() tea.Cmd {
	return nil
}

func (m grid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			m.symbol = "^"
			if m.y > 0 {
				m.y--
			}
		case "down":
			m.symbol = "v"
			if m.y < GRID_SIZE-1 {
				m.y++
			}
		case "left":
			m.symbol = "<"
			if m.x > 0 {
				m.x--
			}
		case "right":
			m.symbol = ">"
			if m.x < GRID_SIZE-1 {
				m.x++
			}
		}
	}

	return m, nil
}

func (m grid) View() string {
	s := ""

	for i := 0; i < GRID_SIZE+2; i++ {
		s += "#"
	}
	s += "\n"

	for i := 0; i < GRID_SIZE; i++ {
		s += "#"
		for j := 0; j < GRID_SIZE; j++ {
			if i == m.y && j == m.x {
				s += m.symbol
			} else {
				s += " "
			}
		}
		s += "#\n"
	}

	for i := 0; i < GRID_SIZE+2; i++ {
		s += "#"
	}
	s += "\n"

	return s
}
