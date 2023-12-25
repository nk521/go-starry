package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

type Styles struct {
	TitleField lipgloss.Style
	InputField lipgloss.Style
}

var (
	commandFieldBlurred       = "Press / to go in command mode!"
	commandFieldFocused       = "Press esc to go in normal mode!"
	commandFieldPromptBlurred = "⭐ "
	commandFieldPromptFocused = "🌟 "
)

func DefaultStyles() *Styles {
	s := new(Styles)
	s.InputField = lipgloss.NewStyle().BorderForeground(lipgloss.Color("#F25D94")).BorderStyle(lipgloss.NormalBorder())
	s.TitleField = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#FFF7DB")).Background(lipgloss.Color("#F25D94"))
	return s
}

type model struct {
	title            string
	commandField     textinput.Model
	command          string
	quitting         bool
	focusedElsewhere bool
	err              error
	width            int
	height           int
	styles           *Styles
	// leftPane ...
	// rightPane ...
}

func (m model) Init() tea.Cmd {
	m.commandField.Cursor.Blink = true
	// return tea.SetWindowTitle("Bubble Tea Example")
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.styles.InputField.Width(percent(m.width-2, 70))
		m.commandField.Width = m.styles.InputField.GetWidth() - len(m.commandField.Prompt) - 1

		m.styles.TitleField.Padding(1, m.width/2, 1, m.width/2).Margin(1, 0, 0, 0)
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			if !m.focusedElsewhere {
				m.quitting = true
				return m, tea.Quit
			}
		case "/":
			if !m.commandField.Focused() {
				m.commandField.Placeholder = commandFieldFocused
				m.commandField.Prompt = commandFieldPromptFocused
				m.commandField.Focus()
				m.focusedElsewhere = true

				return m, cmd
			}
		case "esc":
			if m.commandField.Focused() {
				m.commandField.Placeholder = commandFieldBlurred
				m.commandField.Prompt = commandFieldPromptBlurred
				m.commandField.Blur()
				m.focusedElsewhere = false

				return m, cmd
			}
		case "enter":
			val := m.commandField.Value()
			if len(val) > 0 {
				m.command = val
			}
		}
	}

	m.commandField, cmd = m.commandField.Update(msg)
	return m, cmd
}

func (m model) View() string {
	doc := strings.Builder{}
	if m.width == 0 {
		return "Starting starry... "
	}
	commandFieldRender := m.styles.InputField.Render(m.commandField.View())
	commandFieldPlacement := lipgloss.Place(
		lipgloss.Width(commandFieldRender),
		lipgloss.Height(commandFieldRender),
		lipgloss.Left,
		lipgloss.Bottom,
		commandFieldRender,
	)

	titleRender := m.styles.TitleField.Render("Starry")
	titleRenderHeight, titleRenderWidth := lipgloss.Height(titleRender), lipgloss.Width(titleRender)
	_ = titleRenderHeight
	doc.WriteString(lipgloss.Place(titleRenderWidth, titleRenderHeight, lipgloss.Center, lipgloss.Top, titleRender))

	doc.WriteString(commandFieldPlacement)

	return doc.String()
}

func getCommandInput() *textinput.Model {
	ti := textinput.New()
	ti.Placeholder = commandFieldBlurred
	ti.Prompt = commandFieldPromptBlurred
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F25D94")).Italic(true).Faint(true)
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F25D94"))

	ti.Width = 20

	return &ti
}

func initialModel() model {
	styles := DefaultStyles()

	return model{
		commandField: *getCommandInput(),
		title:        "Starry",
		quitting:     false,
		styles:       styles,
	}
}

func RednerTUI() {

	if _, err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func percent(num, percent int) int {
	return num * percent / 100
}
