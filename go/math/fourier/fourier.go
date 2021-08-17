package fourier

import (
	"math"
	"math/cmplx"
)

func FFT(in []complex128) []complex128 {
	// radix 2 Cooley-Tukey algorithm
	n := len(in)

	if !(n&(n-1) == 0) {
		panic("This algorithm only works when length of n is power of two")
	}

	if n == 1 {
		return in
	}

	// split into even and odd slices:
	inEven, inOdd := splitToEvenAndOdds(in)
	yEven, yOdd := FFT(inEven), FFT(inOdd)

	exponent := complex(0, -2*float64(math.Pi)/float64(n))
	omega := cmplx.Exp(exponent) // primitive nth root of unity

	y := make([]complex128, n)

	for i := 0; i < n/2; i++ {
		y[i] = yEven[i] + cmplx.Pow(omega, complex(float64(i), 0))*yOdd[i]
		y[i+n/2] = yEven[i] - cmplx.Pow(omega, complex(float64(i), 0))*yOdd[i]
	}

	return y
}

// splitToEvenAndOdds splits the slice into even and odd index parts.
// in: [0, 1, 2, 3, 4, 5, 6] -> even: [0, 2, 4, 6], odd: [1, 3, 5]
func splitToEvenAndOdds(in []complex128) (even, odd []complex128) {
	n := len(in)

	even = make([]complex128, n/2+n%2)
	odd = make([]complex128, n/2)

	for i, val := range in {
		if i%2 == 0 { // 0 -> 0, 2 -> 1, 4 -> 2, 6 -> 3
			even[i/2] = val
		} else { // 1 -> 0, 3 -> 1 -> 5 -> 2, 7 -> 3
			odd[(i-1)/2] = val
		}
	}

	return even, odd
}
