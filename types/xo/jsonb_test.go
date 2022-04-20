package xo

import "testing"

// TODO: write proper tests
func TestJSONB(t *testing.T) {
	for _, rawData := range []string{`{"name": "Hello"}`, `["a", "b"]`, `"asd"`} {
		var jsonB Jsonb
		err := jsonB.Scan(rawData)
		if err != nil {
			panic(err)
		}

		println(string(jsonB))

		var nullJsonB NullJsonb
		err = nullJsonB.Scan(rawData)
		if err != nil {
			panic(err)
		}

		println(string(nullJsonB.Jsonb))
	}
}
