# Discrete Fast Fourier Transform

FFT is fast way to calculate the DFT matrix multiplication. It exploits symmetries of roots of unity and thus we get the time complexity from O(n^2) down to O(n log n).

Key Idea:

1. Split polynomial into even and odd terms $A(x) = A_e(x) + A_o(x)$
    - Even terms form an even function: $A_e(x) = A_e(-x)$
    - Odd terms form an odd function: $A_o(x) = -A_o(-x)$
2. Recursively compute $A_e(y) \text{ and } A_o(y) \text{ for } y \in X^2 = \{x^2 | x \in X\}$
3. Combine $A(x) = A_e(x^2) + x \cdot A_o(x^2)$

I like to think about the DFT as a set of dot products between the signal and each base function (so conceptually it is in no way different than any other linear coordinate transform). And the FFT is a way to avoid explicitly computing every single dot product.

## Python implementation

```python
from math import e, pi
from typing import List

import numpy as np


def fft(p: List[complex]):
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


def ifft(p: List[complex]):
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
