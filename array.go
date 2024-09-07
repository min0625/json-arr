package arr

import "encoding/json"

// Array handles empty slice marshal to `[]` instead of `null`.
type Array[T any] []T

func (a Array[T]) MarshalJSON() ([]byte, error) {
	if len(a) == 0 {
		return []byte("[]"), nil
	}

	return json.Marshal([]T(a))
}

func (a *Array[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[]T)(a))
}
