package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostcodeCounter(t *testing.T) {
	postCode := "10200"
	var postCodeBytes postCodeName
	copy(postCodeBytes[:], postCode)
	startTime := int64(10)
	endTime := int64(15)

	t.Run("Post code returns correct str postcode", func(t *testing.T) {
		pcc := NewPostCodeCounter(postCode, startTime, endTime)
		require.Equal(t, postCode, pcc.GetStrPostCode())
	})
	t.Run("Post code not match", func(t *testing.T) {
		{
			pcc := NewPostCodeCounter(postCode, startTime, endTime)
			var wrongPostCodeBytes postCodeName
			copy(wrongPostCodeBytes[:], "10000")
			pcc.Check(wrongPostCodeBytes, startTime, endTime)
			require.Equal(t, pcc.Counter, int64(0))
		}

		{
			pcc := NewPostCodeCounter(postCode, startTime, endTime)
			pcc.Check(postCodeBytes, startTime+1, endTime)
			require.Equal(t, pcc.Counter, int64(0))
		}

		{
			pcc := NewPostCodeCounter(postCode, startTime, endTime)
			pcc.Check(postCodeBytes, startTime, endTime-1)
			require.Equal(t, pcc.Counter, int64(0))
		}
	})

	t.Run("Post code match", func(t *testing.T) {
		{
			pcc := NewPostCodeCounter(postCode, startTime, endTime)
			pcc.Check(postCodeBytes, startTime-1, endTime)
			pcc.Check(postCodeBytes, startTime, endTime+1)
			pcc.Check(postCodeBytes, startTime, endTime+4)
			require.Equal(t, pcc.Counter, int64(3))
		}
		{
			pcc := NewPostCodeCounter(postCode, 1, 3)
			pcc.Check(postCodeBytes, 22, 4)
			require.Equal(t, pcc.Counter, int64(1))
		}
		{
			pcc := NewPostCodeCounter(postCode, 21, 22)
			pcc.Check(postCodeBytes, 18, 2)
			require.Equal(t, pcc.Counter, int64(1))
		}
	})
}
