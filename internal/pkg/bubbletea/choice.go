package bubbletea

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strings"
)

type TeaChoiceModel struct {
	choices []string // items on the to-do list
	Choice  string
	cursor  int    // which to-do list item our Cursor is pointing at
	header  string // which to-do items are Selected
}

func NewChoiceTea(choices []string, header string) *TeaChoiceModel {
	return &TeaChoiceModel{
		// Our to-do list is a grocery list
		choices: choices,
		header:  header,
	}
}

func (m TeaChoiceModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m TeaChoiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch _msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		switch _msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		// The "up" and "k" keys move the Cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		// The "down" and "j" keys move the Cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		// The "enter" key and the spacebar (a literal space) toggle
		// the Selected state for the item that the Cursor is pointing at.
		case "enter", " ":
			m.Choice = m.choices[m.cursor]
			return m, tea.Quit
		}
	}
	// Return the updated TeaListModel to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m TeaChoiceModel) View() string {
	s := strings.Builder{}
	s.WriteString(m.header)

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func GetChoiceTeaResult(header string, choices []string) string {
	ct := NewChoiceTea(choices, header)
	p := tea.NewProgram(ct)

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(TeaChoiceModel); ok && m.Choice != "" {
		fmt.Printf("\n---\nYou chose %s!\n", m.Choice)
		return m.Choice
	}
	return ""
}
