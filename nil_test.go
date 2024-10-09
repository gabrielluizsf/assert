package assert

import (
	"slices"
	"testing"
)

type creator interface {
	create()
}
type accountCreator struct{}

func (c accountCreator) create() {}

type socialNetwork struct {
	creator
}

func TestNil(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "should be nil"
	sn := socialNetwork{
		creator: nil,
	}
	Nil(t, sn.creator, fatalMessage)
	sn.creator = new(accountCreator)
	Nil(spy, sn.creator, fatalMessage)
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
}
