package youtube_music

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	config "github.com/nk521/go-starry/config"
)

func GetHeaders() {
	p := tea.NewProgram(getHeaderModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
}

func getHeaderModel() model {
	ti := textinput.New()
	ti.Placeholder = "Paste your headers here ..."
	curr_headers := config.GetConfig().Login.Headers
	if len(curr_headers) > 0 {
		ti.Placeholder = "You already have headers -> " + curr_headers
	}
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#D1345B")).Italic(true).Faint(true)
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#D1345B"))
	ti.Focus()
	ti.CharLimit = 0

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// m.textInput.SetWidth(msg.Width)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlD:
			headers := m.textInput.Value()
			if len(headers) > 0 {
				config.SetRawConfig("login.headers", headers)
				err := config.SaveRawConfig()
				if err != nil {
					log.Panicln(err)
				}
			}
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	new := lipgloss.NewStyle().Foreground(lipgloss.Color("#D1345B")).Italic(true)
	return fmt.Sprintf(
		"Please paste the request headers from Firefox ...\n\n%s\n\n%s",
		m.textInput.View(),
		"(Press "+new.Render("Esc")+" to quit! (No changes will be done to config))",
	) + "\n"
}
