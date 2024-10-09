package assert

import (
	"slices"
	"testing"
)

func TestEqual(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "should be equal"
	Equal(t, "10", "10", fatalMessage)
	Equal(spy, "10", "7", fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}

func TestEqualBytes(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "should be equal"
	msg := "Hello World"
	b := []byte(msg)
	Equal(t, b, b, fatalMessage)
	Equal(spy, b, append(b, '7'), fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}

func TestEqualSlices(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "should be equal"
	msg := "Hello World"
	s := []string{msg}
	Equal(t, s, s, fatalMessage)
	Equal(spy, s, append(s, "7"), fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}
