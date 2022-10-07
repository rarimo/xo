package xo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: write proper tests
func TestJSONB(t *testing.T) {
	t.Run("Bytes slice can handle modification of original data", func(t *testing.T) {
		inputs := [][]byte{[]byte(`{"name": "Hello"}`), []byte(`["a", "b"]`), []byte(`"asd"`)}
		expectedOutputs := make([][]byte, len(inputs))
		for i, input := range inputs {
			expectedOutputs[i] = make([]byte, len(input))
			copy(expectedOutputs[i], input)
		}

		jsonBs := make([]Jsonb, len(inputs))
		for i, input := range inputs {
			err := jsonBs[i].Scan(input)
			if err != nil {
				panic(err)
			}

			input[0] = 12 // modify original slice in some way
		}

		for i, jsonB := range jsonBs {
			assert.Equal(t, string(expectedOutputs[i]), string(jsonB))
		}
	})
	t.Run("Bytes slice can handle modification of original data for NullJsonB", func(t *testing.T) {
		inputs := [][]byte{[]byte(`{"name": "Hello"}`), []byte(`["a", "b"]`), []byte(`"asd"`)}
		expectedOutputs := make([][]byte, len(inputs))
		for i, input := range inputs {
			expectedOutputs[i] = make([]byte, len(input))
			copy(expectedOutputs[i], input)
		}

		jsonBs := make([]NullJsonb, len(inputs))
		for i, input := range inputs {
			err := jsonBs[i].Scan(input)
			if err != nil {
				panic(err)
			}

			input[0] = 12 // modify original slice in some way
		}

		for i, jsonB := range jsonBs {
			assert.Equal(t, string(expectedOutputs[i]), string(jsonB.Jsonb))
		}
	})
}
