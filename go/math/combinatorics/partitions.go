package combinatorics

type twoInts struct {
	k, n int
}

type PartitionSolver struct {
	cache map[twoInts]int
}

func NewPartitionSolver() *PartitionSolver {
	return &PartitionSolver{cache: make(map[twoInts]int)}
}

// Partitions returns the number of partitions an integer has.
// https://en.wikipedia.org/wiki/Partition_function_(number_theory)
func (ps *PartitionSolver) Partitions(n int) int {
	if n <= 1 {
		return 1
	}

	res := 0

	// n can be partitioned into exactly 1, 2, 3, ..., n parts
	// so the result is p(n, 1) + p(n, 2) + p(n, 3) + ... + p(n, n)
	for i := 1; i <= n; i++ {
		res += ps.partition(n, i)
	}

	return res
}

// partition returns how many ways n can be partitioned into exactly k parts
// https://math.stackexchange.com/a/1344553
// https://www.mathpages.com/home/kmath091.htm
// recurrence relation
// p(n, k) = p(n - 1, k - 1) + p(n -k, k)
func (ps *PartitionSolver) partition(n, k int) int {
	// one number or all ones, e.g. 4 or 1+1+1+1
	if k == 1 || k == n {
		return 1
	}

	// e.g. n = 3, k = 4: how many ways we can partition 3 into 4 numbers
	if n < k {
		return 0
	}

	nk := twoInts{n, k}

	// already in cache
	if val, ok := ps.cache[nk]; ok {
		return val
	}

	// not in cache -> calculate using recurrence and add to cache
	val := ps.partition(n-1, k-1) + ps.partition(n-k, k)

	ps.cache[nk] = val

	return val
}
