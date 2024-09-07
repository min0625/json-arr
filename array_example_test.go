package arr_test

import (
	"encoding/json"
	"fmt"

	jarr "github.com/min0625/json-arr"
)

func Example() {
	type User struct {
		Name      string             `json:"name"`
		Followers jarr.Array[string] `json:"followers"`
	}

	users := []*User{
		{
			Name: "Alice",
			Followers: []string{
				"Bob",
			},
		},
		{
			Name:      "Bob",
			Followers: nil,
		},
	}

	data, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// Output:
	// [{"name":"Alice","followers":["Bob"]},{"name":"Bob","followers":[]}]
}
