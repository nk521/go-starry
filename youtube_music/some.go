package youtube_music

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	config "github.com/nk521/go-starry/config"
)

func GetCookie() {
	p := tea.NewProgram(getCookieModel())
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

func getCookieModel() model {
	ti := textinput.New()
	ti.Placeholder = "Paste your cookies here ..."
	curr_cookies := config.GetConfig().Login.Cookies
	if len(curr_cookies) > 0 {
		ti.Placeholder = "You already have cookies -> " + curr_cookies
	}
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#D1345B")).Italic(true).Faint(true)
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#D1345B"))
	ti.Focus()

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
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			cookie := m.textInput.Value()
			if len(cookie) > 0 {
				config.SetRawConfig("login.cookies", cookie)
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
		"Follow the given steps ...\n"+
			"1. Install "+new.Render("cookie-editor (https://github.com/Moustachauve/cookie-editor)")+" plugin on your browser.\n"+
			"2. Open a logged in session of Youtube Music.\n"+
			"3. Click on the plugin and then click on 'Export' on bottom right.\n"+
			"4. Click on 'Header String'.\n"+
			"5. This will copy the Cookies to your clipboard.\n"+
			"6. Paste that here ...\n\n%s\n\n%s",
		m.textInput.View(),
		"(Press "+new.Render("Esc")+" to quit!)",
	) + "\n"
}
