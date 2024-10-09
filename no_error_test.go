package assert

import (
	"errors"
	"slices"
	"testing"
)

func TestNoError(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "should be nil"
	NoError(t, nil, fatalMessage)
	NoError(spy, errors.New("error"), fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}
