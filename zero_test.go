package assert

import (
	"slices"
	"testing"
)

func TestZero(t *testing.T) {
	fatalMessage := "should be zero value"
	Zero(t, "", fatalMessage)
	spy := newSpy(t)
	Zero(spy, "10", fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}
