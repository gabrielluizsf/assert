package assert

import (
	"reflect"

	"github.com/i9si-sistemas/stringx"
)

// typeOf returns the name of the concrete type of the given value.
// For pointers, it dereferences and returns the base type name.
// Returns an empty string for nil values or unnamed types.
func typeOf(data any) string {
	empty := stringx.Empty.String()
	if data == nil {
		return empty
	}

	t := reflect.TypeOf(data)
	if t == nil {
		return empty
	}

	// Dereference pointers
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Return the type name (empty for unnamed types)
	return t.Name()
}
