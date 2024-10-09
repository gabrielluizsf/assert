package assert

import (
	"reflect"
)

// Zero checks if the provided value is equal to the zero value of its type.
// If the value is not equal to the zero value, it reports a fatal error
// using the provided fatalMessage.
//
//	var myVar string
//	assert.Zero(t, myVar, "myVar should be empty string")
func Zero(t T, value any, fatalMessage string) {
	tester := initTest(t)
	zeroValue := reflect.Zero(reflect.TypeOf(value))
	configureTest(tester, value, zeroValue)
	if !reflect.DeepEqual(value, zeroValue.Interface()) {
		tester.Fatal(fatalMessage)
	}
}
