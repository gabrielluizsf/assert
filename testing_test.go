package assert

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/i9si-sistemas/assert/call"
)

func TestTesting(t *testing.T) {
	ctx := context.Background()
	test := newTest(ctx, t)
	if test.Context() != ctx {
		t.Fail()
	}
	caller := call.Caller()
	test.setCaller(caller)
	if len(test.Caller) == 0 {
		t.Fail()
	}
	if test.Caller != caller {
		t.Fatal("invalid caller")
	}
	spy := newSpy(initTest(test))
	test = newTest(ctx, initTest(spy))
	var called bool
	test.Cleanup(func() {
		called = true
	})
	if !called {
		t.Fail()
	}
	test.Fail()
	if !test.Failed() {
		t.Fail()
	}
	test.Helper()
	if !spy.calls["Helper"] {
		t.Fail()
	}
	funcName := "Setenv"
	test.Skip()
	if !test.Skipped() {
		t.Fail()
	}
	key, value := "gopher", "batata"
	test.Setenv(key, value)
	if !spy.calls[funcName] {
		t.Fail()
	}
	if !slices.Equal(spy.params[funcName].Args, []any{key, value}) {
		t.Fail()
	}
	dir := test.TempDir()
	if dir != "./spy" {
		t.Fail()
	}
	if !spy.calls["TempDir"] {
		t.Fail()
	}
	var code int
	test.exit = func(c int) {
		code = c
	}
	test.Fatal()
	if code != 1 {
		t.Fail()
	}
	msg := test.failedMessage()
	Equal(t, msg, "\n{}\n\n", "should be equal")
	failedMessage := "Hello World"
	msg = test.failedMessage(failedMessage)
	expected := fmt.Sprintf("\n{\n  \"message\": \"%s\"\n}\n\n", failedMessage)
	Equal(t, msg, expected, "should be equal")
	test.result = msg
	test.expected = expected
	msg = test.failedMessage(failedMessage)
	expected = fmt.Sprintf("\n{\n  \"result\": \"\\n{\\n  \\\"message\\\": \\\"%s\\\"\\n}\\n\\n\",\n  \"expected\": \"\\n{\\n  \\\"message\\\": \\\"%s\\\"\\n}\\n\\n\",\n  \"message\": \"%s\"\n}\n\n", failedMessage, failedMessage, failedMessage)
	Equal(t, msg, expected, "should be equal")
	jsonMarshalIndent = func(v any, prefix, indent string) ([]byte, error) {
		return nil, errors.New("marshal error")
	}
	msg = test.failedMessage()
	Zero(t, msg, "should be void string")
	jsonMarshalIndent = json.MarshalIndent
}
