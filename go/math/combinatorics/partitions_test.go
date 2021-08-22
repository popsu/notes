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
		{in: 1000, wantString: "24_061_467_864_032_622_473_692_149_727_991"},
		{in: 10000, wantString: "36167251325636293988820471890953695495016030339315650422081868605887952568754066420592310556052906916435144"},
	}
	for _, tt := range testCases {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			solver := NewPartitionSolver()
			got := solver.Partitions(tt.in)

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
