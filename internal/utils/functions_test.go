package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumFloat32(t *testing.T) {
	// описываем массив структур с тест кейсами
	testCases := []struct {
		name     string
		args     []float32
		expected float32
	}{
		{
			name:     "array empty",
			args:     []float32{},
			expected: 0, // можем не писать это, т.к. у любого числа пустое значение будет 0
		},
		{
			name: "array nil",
			args: []float32{},
			// например здесь мы не пишем expected: 0
		},
		{
			name:     "success",
			args:     []float32{3.33, 5, 0.67, 1},
			expected: 10,
		},
		{
			name:     "negative result",
			args:     []float32{-5, -11, 4},
			expected: -12,
		},
	}

	// бежим в цикле по тест кейсам
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			actual := SumFloat32(tt.args...)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
