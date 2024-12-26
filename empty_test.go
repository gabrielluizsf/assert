package assert

import (
	"slices"
	"testing"
)

func TestEmpty(t *testing.T) {
	fatalMessage := "should be zero value"
	Empty(t, "", fatalMessage)
	spy := newSpy(t)
	Empty(spy, "10", fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}

	Empty(t, 0, fatalMessage)
	spy = newSpy(t)
	Empty(spy, 10, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}

	var user struct{
		Name string
	}
	Empty(t, user, fatalMessage)
	spy = newSpy(t)
	user = struct{Name string}{Name: "gabrielluizsf"}
	Empty(spy, user, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}
