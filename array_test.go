package arr_test

import (
	"encoding/json"
	"testing"

	jarr "github.com/min0625/json-arr"
	"github.com/stretchr/testify/assert"
)

func Test_Array_MarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		give jarr.Array[int]
		want string
	}{
		{
			give: []int{1, 2, 3},
			want: "[1,2,3]",
		},
		{
			give: []int{},
			want: "[]",
		},
		{
			give: nil,
			want: "[]",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(tt.give)
			assert.NoError(t, err)

			assert.JSONEq(t, tt.want, string(got))
		})
	}
}

func Test_Array_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		give string
		want jarr.Array[int]
	}{
		{
			give: "[1,2,3]",
			want: []int{1, 2, 3},
		},
		{
			give: "[]",
			want: []int{},
		},
		{
			give: "null",
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.give, func(t *testing.T) {
			t.Parallel()

			var got jarr.Array[int]

			err := json.Unmarshal([]byte(tt.give), &got)
			assert.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}
