package pkg

import "encoding/json"

func JsonMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func JsonUnmarshal(data []byte, v any) error {
	return json.Unmarshal(data, &v)
}
