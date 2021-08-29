# Cyclotomic Polynomials

For $n \in Z_+$ the $n$th cyclotomic polynomial is the unique irreducible polynomial whose roots are all $n$th **primitive** roots of unity:

$\displaystyle \Phi_{n}(x) = \prod_{\stackrel{1\leq k\leq n}{\gcd(k,n)=1}} \left(x-e^{2i\pi {\frac{k}{n}}}\right)$.

$x$ is a root of $x^{n} -1$ if and only if it is a $d$th primitive root of unity for some $d$ that divides $n$:

$\displaystyle \prod_{d \mid n} \Phi_{d}(x) = x^{n} - 1$

The coefficients of cyclotomic polynomials are always integers. The first cyclotomic polynomial with coefficient other than $1, -1$ or $0$ is when $n = 105$.

??? example "Example - $\Phi_{1}(x)$"
    The only primitive 1st root of unity is 1. Thus $\Phi_{1}(x) = x - 1$

??? example "Example - $\Phi_{2}(x)$"
    The only primitive 2nd root of unity is -1. Thus $\Phi_{2}(x) = x - (-1) = x + 1$

??? example "Example - $\Phi_{3}(x)$"
    The 2 primitive 3rd roots of unity are $e^{2i \pi \cdot \frac{1}{3}} = - \frac{1}{2} + i \frac{\sqrt{3}}{2}$ and $e^{2i \pi \cdot \frac{2}{3}} = - \frac{1}{2} - i \frac{\sqrt{3}}{2}$

    Thus $\Phi_{3}(x) = (x - (-\frac{1}{2} + i \frac{\sqrt{3}}{2})) \cdot (x - (-\frac{1}{2} - i \frac{\sqrt{3}}{2})) = (x + \frac{1}{2} - i \frac{\sqrt{3}}{2}) \cdot (x + \frac{1}{2} + i \frac{\sqrt{3}}{2})$

    $= x^2 + \frac{1}{2}x + \cancel{i\frac{\sqrt{3}}{2}x} + \frac{1}{2}x + \frac{1}{4} + \cancel{i\frac{\sqrt{3}}{4}} - \cancel{i\frac{\sqrt{3}}{2}x} - \cancel{i\frac{\sqrt{3}}{4}} - (i\frac{\sqrt{3}}{2})^2$

    $= x^2 + x + \frac{1}{4} - (i^{2}\frac{3}{4}) = x^2 + x + \frac{1}{4} - (- \frac{3}{4}) = x^2 + x + 1$

??? example "Example - $\Phi_{6}(x)$"
    The 2 primitive 6th roots of unity are $e^{2i \pi \cdot \frac{1}{6}} = \frac{1}{2} + i \frac{\sqrt{3}}{2}$ and $e^{2i \pi \cdot \frac{5}{6}} = \frac{1}{2} - i \frac{\sqrt{3}}{2}$

    Thus $\Phi_{6}(x) = (x - (\frac{1}{2} + i \frac{\sqrt{3}}{2})) \cdot (x - (\frac{1}{2} - i \frac{\sqrt{3}}{2})) = (x - \frac{1}{2} - i \frac{\sqrt{3}}{2}) \cdot (x - \frac{1}{2} + i \frac{\sqrt{3}}{2})$

    $= x^2 - \frac{1}{2}x + \cancel{i\frac{\sqrt{3}}{2}x} - \frac{1}{2}x + \frac{1}{4} - \cancel{i \frac{\sqrt{3}}{4}} - \cancel{i\frac{\sqrt{3}}{2}x} + \cancel{i\frac{\sqrt{3}}{4}} - (i\frac{\sqrt{3}}{2})^2$

    $= x^2 - x + \frac{1}{4} - (i^{2}\frac{3}{4}) = x^2 - x + \frac{1}{4} - (- \frac{3}{4}) = x^2 - x + 1$

??? example "Example $\displaystyle \prod_{d \mid 6} \Phi_{d}(x)$"
    $\displaystyle \prod_{d \mid 6} \Phi_{d}(x) = \Phi_{1}(x) \cdot \Phi_{2}(x) \cdot \Phi_{3}(x) \cdot \Phi_{6}(x) = (x - 1) \cdot (x + 1) \cdot (x^2 + x + 1) \cdot (x^2 - x + 1)$

    $= (x - 1) \cdot (x + 1) \cdot (x^4 - \cancel{x^3} + \cancel{x^2} + \cancel{x^3} - \cancel{x^2} + \cancel{x} + x^2 - \cancel{x} + 1)$

    $= (x - 1) \cdot (x + 1) \cdot (x^4 + x^2 + 1)$

    $= (x - 1) \cdot (x^5 + x^3 + x + x^4 + x^2 + 1) = (x-1) \cdot (x^5 + x^4 + x^3 + x^2 + x + 1)$

    $= (x^6 + \cancel{x^5} + \cancel{x^4} + \cancel{x^3} + \cancel{x^2} + \cancel{x} - \cancel{x^5} - \cancel{x^4} - \cancel{x^3} - \cancel{x^2} - \cancel{x} - 1)$

    $= x^6 - 1$

## Links

- [Cyclotomic polynomial - wikipedia](https://en.wikipedia.org/wiki/Cyclotomic_polynomial)
