package fourier

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFFT(t *testing.T) {
	const epsilon = 1e-8

	testCases := []struct {
		desc string
		in   []complex128
		want []complex128
	}{
		{
			desc: "2 length",
			in:   []complex128{1, 2},
			want: []complex128{3, -1},
		},
		{
			desc: "4 length",
			in:   []complex128{1, 2, 3, 4},
			want: []complex128{
				10,
				-2 + 2i,
				-2,
				-2 - 2i,
			},
		},
		{
			desc: "8 length",
			in:   []complex128{8, 0, 7, 0, 0, 2, 0, 3},
			want: []complex128{
				20,
				8.70710678 - 3.46446609e+00i,
				1 + 1i,
				7.29289322 + 1.05355339e+01i,
				10 - 2.08189956e-15i,
				7.29289322 - 1.05355339e+01i,
				1 - 1i,
				8.70710678 + 3.46446609e+00i,
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			got := FFT(tt.in)
			assert.True(t, equalApprox(tt.want, got, epsilon))
		})
	}
}

func TestIFFT(t *testing.T) {
	const epsilon = 1e-8

	testCases := []struct {
		desc string
		in   []complex128
		want []complex128
	}{
		{
			desc: "2 length",
			in:   []complex128{3, -1},
			want: []complex128{1, 2},
		},
		{
			desc: "4 length",
			in: []complex128{
				10,
				-2 + 2i,
				-2,
				-2 - 2i,
			},
			want: []complex128{1, 2, 3, 4},
		},
		{
			desc: "8 length",
			in: []complex128{
				20,
				8.70710678 - 3.46446609e+00i,
				1 + 1i,
				7.29289322 + 1.05355339e+01i,
				10 - 2.08189956e-15i,
				7.29289322 - 1.05355339e+01i,
				1 - 1i,
				8.70710678 + 3.46446609e+00i,
			},
			want: []complex128{8, 0, 7, 0, 0, 2, 0, 3},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			got := IFFT(tt.in)
			assert.True(t, equalApprox(tt.want, got, epsilon))
		})
	}
}

func TestSplitToEvenAndOdds(t *testing.T) {
	testCases := []struct {
		desc     string
		in       []complex128
		wantEven []complex128
		wantOdd  []complex128
	}{
		{
			desc:     "1 length",
			in:       []complex128{0},
			wantEven: []complex128{0},
			wantOdd:  []complex128{},
		},
		{
			desc:     "2 length",
			in:       []complex128{0, 1},
			wantEven: []complex128{0},
			wantOdd:  []complex128{1},
		},
		{
			desc:     "4 length",
			in:       []complex128{0, 1, 2, 3},
			wantEven: []complex128{0, 2},
			wantOdd:  []complex128{1, 3},
		},
		{
			desc:     "5 length",
			in:       []complex128{0, 1, 2, 3, 4},
			wantEven: []complex128{0, 2, 4},
			wantOdd:  []complex128{1, 3},
		},
		{
			desc:     "8 length",
			in:       []complex128{0, 1, 2, 3, 4, 5, 6, 7},
			wantEven: []complex128{0, 2, 4, 6},
			wantOdd:  []complex128{1, 3, 5, 7},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			gotEven, gotOdd := splitToEvenAndOdds(tt.in)

			assert.Equal(t, tt.wantEven, gotEven)
			assert.Equal(t, tt.wantOdd, gotOdd)
		})
	}
}

// equalApprox returns false iff any of the differences
// of the real or imaginary parts is equal or over epsilon
func equalApprox(a, b []complex128, epsilon float64) bool {
	n := len(a)
	if len(b) != n {
		return false
	}

	ar := make([]float64, n)
	ai := make([]float64, n)
	br := make([]float64, n)
	bi := make([]float64, n)

	for i := 0; i < n; i++ {
		ar[i] = real(a[i])
		ai[i] = imag(a[i])
		br[i] = real(b[i])
		bi[i] = imag(b[i])
	}

	for i := 0; i < n; i++ {
		if math.Abs((ar[i] - br[i])) >= epsilon {
			return false
		}
		if math.Abs((ai[i] - bi[i])) >= epsilon {
			return false
		}
	}

	return true
}
