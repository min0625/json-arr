# json-arr
[![Go Reference](https://pkg.go.dev/badge/github.com/min0625/json-arr.svg)](https://pkg.go.dev/github.com/min0625/json-arr)

A JSON array type uses an empty array instead of null when marshaling to JSON.

## Installation
```sh
go get github.com/min0625/json-arr
```

## Example
```go

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

```
