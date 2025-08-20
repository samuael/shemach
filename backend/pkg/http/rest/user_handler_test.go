package rest

import (
	"testing"

	"github.com/samuael/shemach/backend/pkg/constants/types"

	goJson "github.com/goccy/go-json"
)

var data = `{
    "phone": "123131231322",
    "password": "abc",
    "confirm_password": "abc"
}`
var userdata *types.UserAuthData

func BenchmarkUnmarshal(b *testing.B) {
	// for range b. {
	// err := json.Unmarshal([]byte(data), &userdata)
	// if err != nil {
	// 	b.Fatal(err)
	// }
	// }
}

func BenchmarkUnmarshalGoJson(b *testing.B) {
	// for range b.N {
	err := goJson.Unmarshal([]byte(data), &userdata)
	if err != nil {
		b.Fatal(err)
	}
	// }
}
