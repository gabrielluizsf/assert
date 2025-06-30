package assert

import (
	"testing"

	"github.com/i9si-sistemas/stringx"
)

func TestTypeOf(t *testing.T) {
	type (
		Billing struct {
			Credits int `json:"credits"`
		}
		Account struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Billing `json:"billing"`
		}
		User struct {
			Account `json:"account"`
		}
	)

	var user User

	typeStr := typeOf(user)
	Equal(t, typeStr, "User")

	typeStr = typeOf(&user)
	Equal(t, typeStr, "User")

	typeStr = typeOf(0)
	Equal(t, typeStr, "int")

	typeStr = typeOf(nil)
	Equal(t, typeStr, stringx.Empty)

	typeStr = typeOf(func() {})
	Equal(t, typeStr, stringx.Empty)

	type Handler func(v any) error

	typeStr = typeOf(Handler(nil))
	Equal(t, typeStr, "Handler")

	typeStr = typeOf(func(v any) error { return nil })
	Equal(t, typeStr, stringx.Empty)
}
