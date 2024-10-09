package assert

import (
	"errors"
	"slices"
	"testing"
)

func TestError(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "should be not nil"
	Error(t, errors.New("error"), fatalMessage)
	Error(spy, nil, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}
