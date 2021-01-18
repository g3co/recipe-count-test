package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConvertTime24IntTo12Str(t *testing.T) {
	t.Run("Correct time", func(t *testing.T) {
		{
			time, err := convertTime24IntTo12Str(10)
			require.Nil(t, err)
			require.Equal(t, time, "10AM")
		}
		{
			time, err := convertTime24IntTo12Str(12)
			require.Nil(t, err)
			require.Equal(t, time, "12PM")
		}
		{
			time, err := convertTime24IntTo12Str(0)
			require.Nil(t, err)
			require.Equal(t, time, "12AM")
		}
	})

	t.Run("Wrong time format", func(t *testing.T) {
		_, err := convertTime24IntTo12Str(25)
		require.Error(t, err)
	})
}

func TestConvertTimeByteToInt(t *testing.T) {
	t.Run("Correct time", func(t *testing.T) {
		{
			time := workTime{}
			copy(time[:], "10AM")
			intTime, err := convertTimeByteToInt(time)
			require.Nil(t, err)
			require.Equal(t, intTime, int64(10))
		}
		{
			time := workTime{}
			copy(time[:], "10PM")
			intTime, err := convertTimeByteToInt(time)
			require.Nil(t, err)
			require.Equal(t, intTime, int64(22))
		}
		{
			time := workTime{}
			copy(time[:], "12PM")
			intTime, err := convertTimeByteToInt(time)
			require.Nil(t, err)
			require.Equal(t, intTime, int64(12))
		}
		{
			time := workTime{}
			copy(time[:], "12AM")
			intTime, err := convertTimeByteToInt(time)
			require.Nil(t, err)
			require.Equal(t, intTime, int64(0))
		}
	})

	t.Run("Wrong time format", func(t *testing.T) {
		time := workTime{}
		copy(time[:], "32AM")
		_, err := convertTimeByteToInt(time)
		require.Error(t, err)
	})
}
