package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	items        []mydata
	cursor       int
	instructions string
}

type mydata struct {
	name    string
	enabled bool
}

func initialModel() model {
	return model{
		items:  []mydata{{name: "test"}, {name: "test2"}, {name: "test3", enabled: true}, {name: "test4"}, {name: "test5"}},
		cursor: 2,
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

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case "enter":
			m.instructions = m.items[m.cursor].name
			m.items[m.cursor].enabled = !m.items[m.cursor].enabled
		}
	}
	return m, nil
}

func (m model) View() string {
	bkt := "List:\n"
	for i, item := range m.items {
		if m.cursor == i {
			checkbox := "[ ]"
			if m.items[i].enabled == true {
				checkbox = "[x]"
			} else {
				ckeckbox = "[ ]"
			}

			bkt = bkt + item.name + checkbox + " *\n"
		} else {
			bkt = bkt + item.name + "\n"
		}
	}
	bkt = bkt + "\nPress q to quit.\n" + m.instructions
	return bkt
}

func main() {
	prg := tea.NewProgram(initialModel())

	if _, err := prg.Run(); err != nil {
		fmt.Printf("Oopsie...: %v", err)
	}
}
