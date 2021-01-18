package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser(t *testing.T) {
	var buff buffType
	for i := 0; i < fieldsQuantity; i++ {
		buff[i] = make([]byte, 0, bufferSize)
	}

	t.Run("Correct data", func(t *testing.T) {
		buff[0] = append(buff[0][:0], []byte(`  "postcode": "10190",`)...)
		buff[1] = append(buff[1][:0], []byte(`  "recipe": "Creamy Shrimp Tagliatelle And other text",`)...)
		buff[2] = append(buff[2][:0], []byte(`  "delivery": "Saturday 2AM - 10PM"`)...)

		res, err := ParseBuffer(buff)
		require.Nil(t, err)

		var postCode postCodeName
		copy(postCode[:], "10190")
		require.Equal(t, res.PostCode, postCode)

		var recipe recipeName
		copy(recipe[:], "Creamy Shrimp Tagliatelle And other text")
		require.Equal(t, res.Recipe, recipe)

		require.Equal(t, res.Delivery.End, int64(22))
		require.Equal(t, res.Delivery.Start, int64(2))
	})
}
