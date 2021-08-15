# Discrete fourier transform

Let's consider polynomial:
$A(x) = a_0 \cdot x^0 + a_1 \cdot x^1 + a_2 \cdot x^2 + ... + a_{n-1} \cdot x^{n-1}$

There are 3 different representations for this polynomial:

1. Coefficient vector: $a_0 \cdot x^0 + a_1 \cdot x^1 + ... + a_{n-1} \cdot x^{n-1}$
2. Roots: $r_0, r_1, ..., r_{n-1}$ such that $A = c \cdot (x-r_0) \cdot (x - r_1) \cdots (x - r_{n-1})$
3. Samples of the polynomial: $(x_k, y_k)$ for $\;  k=0, 1, ...n-1$ where $A(x_k) = y_k \; \forall k$ and $x_k$'s are distinct

We can perform 3 different operation on the polynomials:

1. Evaluate one polynomial
2. Add two polynomials
3. Multiply two polynomials

The table for complexity for these operations based on the representation

Operation | Coefficients | Roots | Samples
---------| -----|----- | ------
Evaluate | O(n) | O(n) | O(n^2)
Add      | O(n) | $\infty$  | O(n)
Multiply | O(n^2) | O(n) | O(n)

Discrete fourier transform is the transform between coefficient and sample representations. Using FFT (fast fourier transform) algorithm we can perform this operation in O(n log n) time. The inverse operation also has same time complexity.

The idea is that instead of performing O(n^2) operation, we will use FFT to transform the representation, then do the operation in O(n) time, then transform back with IFFT. This way we can perform all 3 operations in O(n lg n) time or less.

We can calculate the DFT with [DFT Matrix](./dft_matrix.md) or more efficiently with using [FFT](./discrete_fft.md)

## Links

- [Divide & Conquer: FFT MIT OpenCourseWare](https://youtu.be/iTMn0Kt18tg)
- [DSP Lecture 10: The Discrete Fourier Transform](https://www.youtube.com/watch?v=jnxTpxB7HR8)
