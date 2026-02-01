package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	items        []mydata
	cursor       int
	instructions string
	isediting    bool
	textInput    textinput.Model
}

type mydata struct {
	name    string
	enabled bool
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "New item: "
	ti.Focus()
	ti.CharLimit = 64
	ti.Width = 20
	return model{
		items:     []mydata{{name: "test"}, {name: "test2"}, {name: "test3"}, {name: "test4"}, {name: "test5"}},
		cursor:    0,
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !m.isediting {
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "c":
				m.isediting = true
				m.textInput.SetValue("")
				m.instructions = "New Item: "
			}
			if len(m.items) == 0 {
				return m, nil
			}

			switch msg.String() {
			case "up":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down":
				if m.cursor < len(m.items)-1 {
					m.cursor++
				}
			case "enter":
				m.items[m.cursor].enabled = !m.items[m.cursor].enabled
			case "backspace":
				if m.cursor != len(m.items) {
					m.items = append(m.items[:m.cursor], m.items[m.cursor+1:]...)
					if m.cursor >= len(m.items) {
						m.cursor = len(m.items) - 1
					}
				}
			}
		} else {
			switch msg.String() {
			case "enter":
				m.isediting = false
				m.items = append(m.items, mydata{name: m.textInput.Value()})
			}
			var cmd tea.Cmd
			m.textInput, cmd = m.textInput.Update(msg)
			m.instructions = "New Item: " + m.textInput.View()

			return m, cmd
		}
	}
	if !m.isediting {
		if len(m.items) == 0 {
			m.instructions = "No items left!"
			return m, nil
		}
		m.instructions = m.items[m.cursor].name

	}
	return m, nil
}

func (m model) View() string {
	bkt := "List:\n"
	for i, item := range m.items {
		checkbox := "[ ]"
		if m.items[i].enabled {
			checkbox = "[x]"
		}

		if m.cursor == i {
			bkt = bkt + item.name + checkbox + " *\n"
		} else {
			bkt = bkt + item.name + checkbox + "\n"
		}
	}
	bkt = bkt + "\n" + m.instructions + "\nPress Enter to toggle on/off.\nPress q to quit.\n"
	return bkt
}

func main() {
	prg := tea.NewProgram(initialModel())

	if _, err := prg.Run(); err != nil {
		fmt.Printf("Oopsie...: %v", err)
	}
}
