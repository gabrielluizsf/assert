package assert

import (
	"slices"
	"testing"
)

func TestFalse(t *testing.T) {
	fatalMessage := "should be false"
	False(t, false, fatalMessage)
	spy := newSpy(t)
	False(spy, true, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}
