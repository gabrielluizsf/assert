package assert

import (
	"slices"
	"testing"
)

func TestNotEmpty(t *testing.T) {
	fatalMessage := "should be zero value"
	NotEmpty(t, "10", fatalMessage)
	spy := newSpy(t)
	NotEmpty(spy, "", fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}

	NotEmpty(t, 10, fatalMessage)
	spy = newSpy(t)
	NotEmpty(spy, 0, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}

	var user struct {
		Name string
	}
	spy = newSpy(t)
	NotEmpty(spy, user, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
	user = struct{ Name string }{Name: "gabrielluizsf"}
	NotEmpty(t, user, fatalMessage)
}
