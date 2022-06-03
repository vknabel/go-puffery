package nav_test

import (
	"testing"

	"github.com/vknabel/go-puffery/nav"
)

func TestPopCmd(t *testing.T) {
	cmd := nav.Pop()
	returnedMsg := cmd()
	msg, ok := returnedMsg.(nav.PagePopMsg)
	if !ok {
		t.Error("expected PopPageMsg, got", msg)
	}
}

func TestPushCmd(t *testing.T) {
	cmd := nav.Push(testModel{})
	returnedMsg := cmd()
	msg, ok := returnedMsg.(nav.PagePushMsg)
	if !ok {
		t.Error("expected PushPageMsg, got", msg)
	}
	if _, ok := msg.Page.(testModel); !ok {
		t.Error("expected PushPageMsg.Page to be a testModel")
	}
}
