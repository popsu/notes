package combinatorics

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionFunction(t *testing.T) {
	testCases := []struct {
		in   int
		want int
	}{
		{in: 0, want: 1},
		{in: 1, want: 1},
		{in: 2, want: 2},
		{in: 3, want: 3},
		{in: 4, want: 5},
		{in: 5, want: 7},
		{in: 6, want: 11},
		{in: 7, want: 15},
		{in: 8, want: 22},
		{in: 9, want: 30},
		{in: 10, want: 42},
		{in: 11, want: 56},
		{in: 12, want: 77},
		{in: 13, want: 101},
		{in: 14, want: 135},
		{in: 15, want: 176},
		{in: 16, want: 231},
		{in: 17, want: 297},
		{in: 18, want: 385},
		{in: 19, want: 490},
		{in: 20, want: 627},
		{in: 21, want: 792},
		{in: 22, want: 1002},
		{in: 100, want: 190_569_292},
	}
	for _, tt := range testCases {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			solver := NewPartitionSolver()
			got := solver.Partitions(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
