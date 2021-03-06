package nav_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vknabel/go-puffery/nav"
)

func TestNavPageDisplaysPushed(t *testing.T) {
	prog := tea.NewProgram(
		nav.NewPage(initialTestModel("initial")),
		tea.WithoutCatchPanics(),
		tea.WithoutRenderer(),
	)
	go func() {
		prog.Send(nav.PagePushMsg{initialTestModel("pushed")})
		go prog.Quit()
	}()
	model, err := prog.StartReturningModel()
	if err != nil {
		t.Error(err)
	}
	stack, ok := model.(nav.NavPage)
	if !ok {
		t.Error("expected model to be a NavPage, got", model)
	}
	testModel, ok := stack.Top().(testModel)
	if !ok {
		t.Error("expected model to be a testModel, got", model)
	}

	if testModel.id != "pushed" {
		t.Error("expected model id to be 'pushed', got", testModel.id)
	}
}

func TestNavPageDisplaysInitialWhenPushedAndPopped(t *testing.T) {
	prog := tea.NewProgram(
		nav.NewPage(initialTestModel("initial")),
		tea.WithoutCatchPanics(),
		tea.WithoutRenderer(),
	)
	go func() {
		prog.Send(nav.PagePushMsg{initialTestModel("pushed")})
		go func() {
			prog.Send(nav.PagePopMsg{})
			go prog.Quit()
		}()
	}()
	model, err := prog.StartReturningModel()
	if err != nil {
		t.Error(err)
	}
	stack, ok := model.(nav.NavPage)
	if !ok {
		t.Error("expected model to be a NavPage, got", model)
	}
	testModel, ok := stack.Top().(testModel)
	if !ok {
		t.Error("expected model to be a testModel, got", model)
	}

	if testModel.id != "initial" {
		t.Error("expected model id to be 'initial', got", testModel.id)
	}
}
