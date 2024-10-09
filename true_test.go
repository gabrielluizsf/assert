package assert

import (
	"slices"
	"testing"
)

func TestTrue(t *testing.T) {
	fatalMessage := "should be true"
	True(t, true, fatalMessage)
	spy := newSpy(t)
	True(spy, false, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}
