package bubbletea

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type (
	errMsg error
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
		// Cool, what was the actual key pressed?
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

type TeaInputModel struct {
	TextInput textinput.Model
	Err       error
	Question  string
}

func NewInputTea(question string, placeholder string) TeaInputModel {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return TeaInputModel{
		TextInput: ti,
		Err:       nil,
	}
}

func (m TeaInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TeaInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.Err = msg
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m TeaInputModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s", m.Question,
		m.TextInput.View(),
		"(esc to quit)",
	) + "\n"
}
