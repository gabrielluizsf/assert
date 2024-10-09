package assert

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestSpy(t *testing.T) {
	spy := newSpy(t)
	var buf io.Writer
	spy.setOutput = func(w io.Writer) {
		buf = w
	}
	_, cancel := spy.LogOutput()
	v, ok := any(buf).(*bytes.Buffer)
	if !ok {
		t.Fail()
	}
	if v == nil {
		t.Fail()
	}
	cancel()
	if buf != os.Stdout {
		t.Fail()
	}
}
