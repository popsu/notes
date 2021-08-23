package combinatorics

import "math/big"

// Partitions returns the number of partitions for n.
// https://en.wikipedia.org/wiki/Partition_function_(number_theory)
// Uses algorithm from mathologer video https://www.youtube.com/watch?v=iJ8pnCO0nTY
func Partitions(n int) *big.Int {
	values := make([]*big.Int, n+1)
	values[0] = big.NewInt(1)

	gpn := GeneralizedPentagonalNumbers(n)
	var curr, prev *big.Int
	j := 0

	// Calculate partition numbers for numbers up to n
	for i := 1; i <= n; i++ {
		curr, j = big.NewInt(0), 0

		// Add/Sub the previous numbers. Indexes are based on generalized pentagonal numbers
		for {
			idx := gpn[j+1] // first index is not used

			if idx > i {
				break
			}

			prev = values[i-idx]

			switch j % 4 { // the logic whether to add or sub is based on j mod 4
			case 0, 1:
				curr = curr.Add(curr, prev)
			case 2, 3:
				curr = curr.Sub(curr, prev)
			}
			j++
		}

		values[i] = curr
	}

	return values[n]
}

func PentagonalNumber(n int) int {
	return (n * (3*n - 1)) / 2
}

// GeneralizedPentagonalNumbers generates the sequence of Generalized Pentagonal Numbers
// without the first 0 up to size n. They are defined to be Pentagonal numbers using values:
// 0, 1, -1, 2, -2, 3, -3, ... So values being: 0, 1, 2, 5, 7, 12, 15, 22, 26, 35, 40, 51
// https://oeis.org/A001318
// This sequence is used in Partitions algorithm
func GeneralizedPentagonalNumbers(n int) []int {
	m := []int{0}

	curr, val := 0, 0

	for {
		curr++

		val = PentagonalNumber(curr)
		m = append(m, val)
		val = PentagonalNumber(-curr)
		m = append(m, val)

		if val > n {
			break
		}
	}

	return m
}
