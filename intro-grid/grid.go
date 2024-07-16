package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const (
	GRID_SIZE = 10

	INTRO = `# Some introduction

This is a simple example of Markdown rendering with Glamour!
Check out the [other examples](https://github.com/charmbracelet/glamour/tree/master/examples) too.

Bye!
`
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

// Styling
var (
	wallStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4"))
)

func (m grid) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
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

	out, err := glamour.Render(INTRO, "dark")

	if err == nil {
		s += out
	} else {
		s += INTRO
	}

	for i := 0; i < GRID_SIZE+2; i++ {
		s += wallStyle.Render(" ")
	}
	s += "\n"

	for i := 0; i < GRID_SIZE; i++ {
		s += wallStyle.Render(" ")
		for j := 0; j < GRID_SIZE; j++ {
			if i == m.y && j == m.x {
				s += m.symbol
			} else {
				s += " "
			}
		}
		s += wallStyle.Render(" ") + "\n"
	}

	for i := 0; i < GRID_SIZE+2; i++ {
		s += wallStyle.Render(" ")
	}
	s += "\n"

	return s
}
