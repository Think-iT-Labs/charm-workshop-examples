package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Mode int

const (
	InsertMode Mode = iota
	CommandMode
)

type MDim struct {
	textArea textarea.Model
	text     string
	mode     Mode
}

func initMDim() MDim {
	ta := textarea.New()
	ta.ShowLineNumbers = false

	return MDim{
		mode:     CommandMode,
		text:     "",
		textArea: ta,
	}
}

func (m MDim) Init() tea.Cmd {
	return textinput.Blink
}

func (m MDim) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		if m.mode == InsertMode {
			switch msg.Type {
			case tea.KeyCtrlC, tea.KeyEsc:
				m.mode = CommandMode
				m.textArea.Blur()
			}
		} else {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "i":
				m.mode = InsertMode
				m.textArea.Focus()
			}
		}
	}

	m.textArea, cmd = m.textArea.Update(msg)
	m.text = m.textArea.View()
	return m, cmd
}

func (m MDim) View() string {
	left := ""
	left += m.textArea.View() + "\n"

	left += "\n"
	if m.mode == InsertMode {
		left += "--- INSERT ---"
	} else {
		left += "--- COMMAND ---"
	}

	right, err := glamour.Render(m.textArea.Value(), "dark")

	if err != nil {
		right = m.textArea.Value()
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		left,
		right,
	)
}
