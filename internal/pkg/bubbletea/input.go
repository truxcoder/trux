package bubbletea

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type TeaInputModel struct {
	TextInput textinput.Model
	Err       error
	Question  string
	Value     string
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
		Question:  question,
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
		case tea.KeyCtrlC, tea.KeyEsc:
			m.Value = ""
			return m, tea.Quit
		case tea.KeyEnter:
			m.Value = m.TextInput.Value()
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

func GetInputTeaResult(question string, placeholder string) string {
	it := NewInputTea(question, placeholder)
	p := tea.NewProgram(it)
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(TeaInputModel); ok && m.Value != "" {
		//fmt.Printf("m.value-->%s\n", m.Value)
		//fmt.Printf("m.view()-->%s\n", m.TextInput.View())
		//fmt.Printf("m.TextInput.Value()-->%s\n", m.TextInput.Value())
		return m.Value
	}
	return ""
}
