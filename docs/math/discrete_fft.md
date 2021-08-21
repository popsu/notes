# Discrete Fast Fourier Transform

FFT is fast way to calculate the DFT matrix multiplication.

!!! info
    DFT is basically evaluating n points of a (n-1)-degree polynomial, thus naive time complexity is O(n^2). FFT gets this down to O(n log n) by exploiting symmetries of roots of unity and even functions.

Key Idea:

1. Split polynomial into even and odd terms $A(x) = A_e(x) + A_o(x)$
    - Even terms form an even function: $A_e(x) = A_e(-x)$
    - Odd terms we can factor x out and get even function
2. Recursively compute $A_e(y) \text{ and } A_o(y) \text{ for } y \in X^2 = \{x^2 | x \in X\}$
    - Because we are using roots of unity and  we have even functions the number of items in this set stays small enough
    - Base case is when the degree is 0, so $A$ is a constant
3. Combine $A(x) = A_e(x^2) + x \cdot A_o(x^2)$

???+ example

    Let $A(x) = 2x^7 + 3x^6 + 4x^5 + 5x^4 + 6x^3 + 7x^2 + 8x + 9$

    Split A into even and odd terms and substitute $x^2 = y$ and we get two degree 3 polynomials (instead of one degree 7):

    -> even terms $A_e(x) =  3x^6 + 5x^4 + 7x^2  + 9 \implies A_e(y) = 3y^3 + 5y^2 + 7y + 9$

    -> odd terms $A_o(x) = 2x^7 + 4x^5 + 6x^3 + 8x = x (2x^6 + 4x^4 + 6x^2 + 8) \implies A_o(y) = 2y^3 + 4y^2 + 6y + 8$

    Overall this is just $A(x) = 3(x^2)^3 + 5(x^2)^2 + 7(x^2) + 9 + x \cdot (2(x^2)^3 + 4(x^2)^2 + 6(x^2) + 8)$

    Another way to think about the DFT as a set of dot products between the signal and each base function (so conceptually it is in no way different than any other linear coordinate transform). And the FFT is a way to avoid explicitly computing every single dot product.

??? note "Python implementation"
    ```python
    from math import e, pi
    from typing import List

    import numpy as np


    def fft(p: List[complex]) -> List[complex]:
        # radix 2 Cooley-Tukey algorithm
        n = len(p)

        if not n and (not (n & (n - 1))):  # check if n is power of 2
            raise ValueError("This implementation requires arrays with length power of two")

        if n == 1:
            return p

        y_even, y_odd = fft(p[::2]), fft(p[1::2])

        omega = e ** (-2j * pi / n)  # primitive nth root of unity
        y = [0] * n
        for i in range(n // 2):
            y[i] = y_even[i] + omega ** i * y_odd[i]
            y[i + n // 2] = y_even[i] - omega ** i * y_odd[i]
        return y


    def ifft(p: List[complex]) -> List[complex]:
        # Take complex conjugates
        p = [np.conj(x) for x in p]
        # Calculate normal fft
        p = fft(p)
        # Take complex conjugates again and scale by 1/n
        n = len(p)
        return [1 / n * np.conj(x) for x in p]
    ```

## Links

- [Even and odd functions - wikipedia](https://en.wikipedia.org/wiki/Even_and_odd_functions)
- [The Fast Fourier Transform (FFT): Most Ingenious Algorithm Ever? - youtube](https://www.youtube.com/watch?v=h7apO7q16V0)
- [Lecture 8 - Algorithms: The Fast Fourier Transform (FFT) - youtube](https://www.youtube.com/watch?v=2V7XT_iiRRw)
- [26. Complex Matrices; Fast Fourier Transform](https://www.youtube.com/watch?v=M0Sa8fLOajA)
