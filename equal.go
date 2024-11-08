package assert

import (
	"reflect"
)

// Equal checks if the provided result is equal to the expected value.
// If the result is not equal to the expected value, it reports a fatal error
// using the provided fatalMessage.
//
//	result := someFunction()
//	expected := 42
//	assert.Equal(t, result, expected, "Expected result to be 42")
func Equal(t T, result, expected any, args ...any) {
	tester := initTest(t)
	configureTest(tester, result, expected)
	if !equal(result, expected) {
		tester.Fatal(args...)
	}
}


func equal(result, expected any) bool {
	return reflect.DeepEqual(result, expected)
}
