package multitext

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	focusIndex int
	lines      []textinput.Model
}

type TrailingNewlines struct {
	Count int
}

func New(firstLine textinput.Model) Model {
	return Model{
		focusIndex: -1,
		lines:      []textinput.Model{firstLine},
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if m.focusIndex == -1 {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlV:
			// TODO: handle paste
		case tea.KeyEnter:
			if m.focusIndex != -1 {
				m.lines[m.focusIndex].Blur()
			}
			if m.focusIndex == len(m.lines)-1 {
				newLine := textinput.New()
				firstLine := m.lines[0]
				newLine.PromptStyle = firstLine.PromptStyle.MarginLeft(
					firstLine.PromptStyle.GetMarginLeft() + len(firstLine.Prompt),
				)
				newLine.Prompt = strings.Repeat(" ", len(firstLine.Prompt))
				newLine.TextStyle = firstLine.TextStyle
				newLine.PlaceholderStyle = firstLine.PlaceholderStyle
				m.lines = append(m.lines, newLine)
			}
			m.focusIndex += 1
			m.lines[m.focusIndex].Focus()

			if m.focusIndex == len(m.lines)-1 && m.lines[m.focusIndex].Value() == "" {
				emptyLineCount := 0
				for i := len(m.lines) - 1; i > 0; i-- {
					if m.lines[i].Value() == "" {
						emptyLineCount += 1
					} else {
						break
					}
				}
				return m, tea.Batch(
					textinput.Blink,
					func() tea.Msg {
						return TrailingNewlines{Count: emptyLineCount}
					},
				)
			}
			return m, textinput.Blink
		case tea.KeyUp:
			if m.focusIndex == 0 {
				return m, nil
			}
			cursor := m.lines[m.focusIndex].Cursor()
			m.lines[m.focusIndex].Blur()
			m.focusIndex -= 1
			m.lines[m.focusIndex].Focus()
			m.lines[m.focusIndex].SetCursor(cursor)
			return m, textinput.Blink
		case tea.KeyDown:
			if m.focusIndex == len(m.lines)-1 {
				return m, nil
			}
			cursor := m.lines[m.focusIndex].Cursor()
			m.lines[m.focusIndex].Blur()
			m.focusIndex += 1
			m.lines[m.focusIndex].Focus()
			m.lines[m.focusIndex].SetCursor(cursor)
			return m, textinput.Blink
		case tea.KeyBackspace:
			deletedLine := m.lines[m.focusIndex]
			if m.focusIndex != 0 && deletedLine.Cursor() == 0 {
				previousLine := m.lines[m.focusIndex-1]
				previousText := previousLine.Value()
				previousLine.SetValue(previousLine.Value() + deletedLine.Value())
				previousLine.SetCursor(len(previousText))
				previousLine.Focus()
				m.focusIndex -= 1
				m.lines = append(m.lines[:m.focusIndex], m.lines[m.focusIndex+1:]...)
				m.lines[m.focusIndex] = previousLine
				return m, textinput.Blink
			}
		}
	}

	var cmd tea.Cmd
	m.lines[m.focusIndex], cmd = m.lines[m.focusIndex].Update(msg)
	return m, cmd
}

func (m *Model) Focus() tea.Cmd {
	if m.focusIndex == -1 {
		m.focusIndex = 0
		m.lines[m.focusIndex].Focus()
	}
	return nil
}

func (m Model) Focused() bool {
	return m.focusIndex != -1
}

func (m *Model) Blur() tea.Cmd {
	if m.focusIndex != -1 {
		m.lines[m.focusIndex].Blur()
		m.focusIndex = -1
	}
	return nil
}

func (m Model) Value() string {
	lines := []string{}
	for _, line := range m.lines {
		lines = append(lines, line.Value())
	}
	return strings.Join(lines, "\n")
}

func (m Model) View() string {
	lines := []string{}
	for _, line := range m.lines {
		lines = append(lines, line.View())
	}
	return strings.Join(lines, "\n")
}
