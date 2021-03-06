package combinatorics

import (
	"math/big"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionFunction(t *testing.T) {
	testCases := []struct {
		in         int
		want       *big.Int
		wantString string
	}{
		{in: 0, want: big.NewInt(1)},
		{in: 1, want: big.NewInt(1)},
		{in: 2, want: big.NewInt(2)},
		{in: 3, want: big.NewInt(3)},
		{in: 4, want: big.NewInt(5)},
		{in: 5, want: big.NewInt(7)},
		{in: 6, want: big.NewInt(11)},
		{in: 7, want: big.NewInt(15)},
		{in: 8, want: big.NewInt(22)},
		{in: 9, want: big.NewInt(30)},
		{in: 10, want: big.NewInt(42)},
		{in: 11, want: big.NewInt(56)},
		{in: 12, want: big.NewInt(77)},
		{in: 13, want: big.NewInt(101)},
		{in: 14, want: big.NewInt(135)},
		{in: 15, want: big.NewInt(176)},
		{in: 16, want: big.NewInt(231)},
		{in: 17, want: big.NewInt(297)},
		{in: 18, want: big.NewInt(385)},
		{in: 19, want: big.NewInt(490)},
		{in: 20, want: big.NewInt(627)},
		{in: 21, want: big.NewInt(792)},
		{in: 22, want: big.NewInt(1002)},
		{in: 100, want: big.NewInt(190_569_292)},
		{in: 1_000, wantString: "24_061_467_864_032_622_473_692_149_727_991"},
		{in: 6_666, wantString: "193655306161707661080005073394486091998480950338405932486880600467114423441282418165863"},
		{in: 10_000, wantString: "36167251325636293988820471890953695495016030339315650422081868605887952568754066420592310556052906916435144"},
	}
	for _, tt := range testCases {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			got := Partitions(tt.in)

			if tt.want != nil {
				assert.Equal(t, tt.want, got)
			} else { // use string to set the value since it's too big
				want := &big.Int{}
				want.SetString(strings.Replace(tt.wantString, "_", "", -1), 10)

				assert.Equal(t, want, got)
			}
		})
	}
}

func TestPentagonalNumbers(t *testing.T) {
	testCases := []struct {
		in   int
		want int
	}{
		{in: 1, want: 1},
		{in: -1, want: 2},
		{in: 2, want: 5},
		{in: -2, want: 7},
		{in: 3, want: 12},
		{in: -3, want: 15},
		{in: 4, want: 22},
		{in: -4, want: 26},
		{in: 5, want: 35},
		{in: -5, want: 40},
		{in: 6, want: 51},
		{in: -6, want: 57},
		{in: 7, want: 70},
		{in: -7, want: 77},
		{in: 8, want: 92},
		{in: 9, want: 117},
		{in: 10, want: 145},
		{in: 11, want: 176},
		{in: 12, want: 210},
		{in: 13, want: 247},
		{in: 14, want: 287},
		{in: 15, want: 330},
	}
	for _, tt := range testCases {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			got := PentagonalNumber(tt.in)

			assert.Equal(t, tt.want, got)
		})
	}
}
