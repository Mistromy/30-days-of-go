package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	items  []string
	cursor int
}

func initialModel() model {
	return model{
		items:  []string{"test", "test2"},
		cursor: 30,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// TODO: implement View method
func (m model) View() string {
	return "Test"
}

func main() {
	prg := tea.NewProgram(initialModel())

	if _, err := prg.Run(); err != nil {
		fmt.Printf("Oopsie...: %v", err)
	}
}
