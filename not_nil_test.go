package assert

import (
	"slices"
	"testing"
)

func TestNotNil(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "should be nil"
	NotNil(t, spy, fatalMessage)
	NotNil(spy, nil, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}
