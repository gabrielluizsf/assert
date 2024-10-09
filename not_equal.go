package assert

import (
	"reflect"
)

// NotEqual checks if the two provided values are not equal.
// If the values are equal, it reports a fatal error using the provided faltalMessage.
//
//	a := 5
//	b := 10
//	assert.NotEqual(t, a, b, "a and b should not be equal")
func NotEqual(t T, v1, v2 any, faltalMessage string) {
	tester := initTest(t)
	configureTest(tester, v1, "not equal")
	if reflect.DeepEqual(v1, v2) {
		tester.Fatal(faltalMessage)
	}
}
