package assert

// False checks if the provided boolean value is false.
// If the value is true, it reports a fatal error using the provided fatalMessage.
//
//	result := someFunction()
//	assert.False(t, result, "Expected result to be false")
func False(t T, ok bool, fatalMessage string) {
	tester := initTest(t)
	configureTest(tester, ok, true)
	if ok {
		tester.Fatal(fatalMessage)
	}
}
