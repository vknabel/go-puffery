package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type questionModel struct {
	Prompt           string
	PromptStyle      lipgloss.Style
	defaultAnswer    bool
	PlaceholderStyle lipgloss.Style
	answer           bool
	AnswerStyle      lipgloss.Style
	answered         bool
}

func newQuestion(prompt string) questionModel {
	return questionModel{
		Prompt:           prompt,
		PromptStyle:      lipgloss.NewStyle(),
		PlaceholderStyle: lipgloss.NewStyle().Foreground(colorPlaceholder),
		AnswerStyle:      lipgloss.NewStyle().Foreground(colorLagoonBubbleBlue),
		defaultAnswer:    true,
	}
}

func (q questionModel) View() string {
	content := q.PromptStyle.Render(q.Prompt)
	if q.answered && q.answer {
		content += q.AnswerStyle.Render("Yes")
	} else if q.answered && !q.answer {
		content += q.AnswerStyle.Render("No")
	} else if q.defaultAnswer {
		content += q.PlaceholderStyle.Render("(Y/n)")
	} else {
		content += q.PlaceholderStyle.Render("(y/N)")
	}
	return content
}

func (q questionModel) Update(msg tea.Msg) (questionModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "y" {
			q.answer = true
			q.answered = true
			return q, nil
		} else if msg.String() == "n" {
			q.answer = false
			q.answered = true
			return q, nil
		} else if msg.Type == tea.KeyEnter {
			q.answered = true
			q.answer = q.defaultAnswer
			return q, nil
		}
	}
	return q, nil
}
