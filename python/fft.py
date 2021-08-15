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


def test_fft_ifft():
    # Random 7th degree polynomial with coefficients from (-10, 10)
    x = np.random.rand(8) * 20 - 10

    # Verify fft gives same results as numpy's fft
    assert np.allclose(fft(x), np.fft.fft(x))
    assert np.allclose(ifft(x), np.fft.ifft(x))
    assert np.allclose(ifft(fft(x)), x)


if __name__ == "__main__":
    test_fft_ifft()
