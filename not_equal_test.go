package assert

import (
	"slices"
	"testing"
)

func TestNotEqual(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "should be equal"
	NotEqual(t, "10", "11", fatalMessage)
	NotEqual(spy, "70", "70", fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}

func TestNotEqualBytes(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "should be equal"
	msg := "Hello World"
	b := []byte(msg)
	NotEqual(spy, b, append(b, '7'), fatalMessage)
	NotEqual(spy, b, b, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}

func TestNotEqualSlices(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "should be equal"
	msg := "Hello World"
	s := []string{msg}
	NotEqual(spy, s, append(s, "7"), fatalMessage)
	NotEqual(spy, s, s, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}
