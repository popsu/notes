# DFT matrix

DFT matrix is a square [Vandermonde matrix](./vandermonde_matrix.md), where the numbers $x_i$ are chosen to be roots of unity.

For example let $n = 8$ and $\omega = e^{\frac{-2\pi i}{8} }= \frac{1}{\sqrt{2}} - \frac{i}{\sqrt{2}}$ (primitive 8th root of unity). We choose the points $x_0, \ldots, x_7$ to be 8th roots of unity:

$$\begin{aligned}
x_0& = \omega ^0 = &1 \\
x_1 &= \omega ^1 = &\omega\\
x_2 &= \omega ^2 = &-i\\
x_3 &= \omega ^3 = &-i \omega\\
x_4 &= \omega ^4 = &-1\\
x_5 &= \omega ^5 = &-\omega\\
x_6 &= \omega ^6 = &i \\
x_7 &= \omega ^7 = &i \omega
\end{aligned}$$

Then the DFT matrix is (note $\omega^8 = 1$):

$${\displaystyle V={\begin{bmatrix}
1&\omega ^{0}&\omega ^{0}&\omega ^{0}&\omega ^{0}&\omega ^{0}&\omega ^{0}&\omega ^{0}\\
1&\omega ^{1}&\omega ^{2}&\omega ^{3}&\omega ^{4}&\omega ^{5}&\omega ^{6}&\omega ^{7}\\
1&\omega ^{2}&\omega ^{4}&\omega ^{6}&\omega ^{8}&\omega ^{10}&\omega ^{12}&\omega ^{14}\\
1&\omega ^{3}&\omega ^{6}&\omega ^{9}&\omega ^{12}&\omega ^{15}&\omega ^{18}&\omega ^{21}\\
1&\omega ^{4}&\omega ^{8}&\omega ^{12}&\omega ^{16}&\omega ^{20}&\omega ^{24}&\omega ^{28}\\
1&\omega ^{5}&\omega ^{10}&\omega ^{15}&\omega ^{20}&\omega ^{25}&\omega ^{30}&\omega ^{35}\\
1&\omega ^{6}&\omega ^{12}&\omega ^{18}&\omega ^{24}&\omega ^{30}&\omega ^{36}&\omega ^{42}\\
1&\omega ^{7}&\omega ^{14}&\omega ^{21}&\omega ^{28}&\omega ^{35}&\omega ^{42}&\omega ^{49}\\
\end{bmatrix}}={\begin{bmatrix}
1&1&1&1&1&1&1&1\\
1&\omega &-i&-i\omega &-1&-\omega &i&i\omega \\
1&-i&-1&i&1&-i&-1&i\\
1&-i\omega &i&\omega &-1&i\omega &-i&-\omega \\
1&-1&1&-1&1&-1&1&-1\\
1&-\omega &-i&i\omega &-1&\omega &i&-i\omega \\
1&i&-1&-i&1&i&-1&-i\\
1&i\omega &i&-\omega &-1&-i\omega &-i&\omega \\
\end{bmatrix}}}$$

## Orthogonality of DFT matrix

All the columns in $V$ are orthogonal and since it's a change-of-base matrix it means the basis is orthogonal. The dot product of two complex vectors is defined as $\textbf{a} \cdot \textbf{b} = \sum a_j \overline b_j$, where $\overline b_j$ is the complex conjugate of $b_j$.

Proof:

Let $\textbf{a}, \textbf{b}$ be two different columns of the DFT matrix and $m, l\; (m\neq l)$ the column numbers of $\textbf{a}$ and $\textbf{b}$.

$$\begin{aligned}
\textbf{a} \cdot \textbf{b} &= \sum_{j=0}^{n-1} a_j \overline b_j &\text{definition of complex dot product}\\
&= \sum_{j=0}^{n-1} \omega^{jm} \omega^{-jl} &\text{definition of $\omega$ in DFT}\\
&= \sum_{j=0}^{n-1} \omega^{(m-l)j} &\text{exponential multiplication rules}\\
&= \sum_{j=0}^{n-1} \omega_{k}^{j} &\omega^{m-l} \text{ is a fixed root of unity, call it }\omega_{k}\\
&= \frac{1 - \omega_{k}^{n}}{1 - \omega_{k}} &\text{geometric series formula, note } \omega_k \neq 1\\
&= \frac{1 - 1}{1 - \omega_k} &{\omega^n = \omega^0 = 1}\\
&= 0\end{aligned}$$

QED.

## Inverse of DFT

Complex matrix with orthonormal columns is called Unitary matrix. Since the DFT columns are orthogonal, we can make the DFT matrix Unitary by scaling the columns to be unit-length:

$$U = \frac{1}{\sqrt{n}} \cdot V \iff \sqrt{n} \cdot U = V$$

For unitary matrixes the inverse is just the conjugate transpose:

$$U^{-1} = \overline U \iff U\overline U = I$$

For invertible matrix $A$ we have: $(kA)^{-1} = k^{-1}A^{-1}$. Using this and the previous statements we can get the inverse of $V$:

$$\begin{aligned}
U^{-1} &= \overline{U} &\iff \\
(\frac{1}{\sqrt{n}} \cdot V)^{-1} &= \overline{(\frac{1}{\sqrt{n}} \cdot V)} &\iff \\
\sqrt{n}V^{-1} &= \frac{1}{\sqrt{n}} \overline{V} &\iff \\
V^{-1} &= \frac{1}{n} \cdot \overline{V}
\end{aligned}$$

So for example the inverse for the DFT with  n = 8, we get:

$${\displaystyle V^{-1} = \frac{1}{n} \cdot \overline V={ \frac{1}{n} \cdot \begin{bmatrix}
1&1&1&1&1&1&1&1\\
1&i\omega &i&-\omega &-1&-i\omega &-i&\omega \\
1&i&-1&-i&1&i&-1&-i\\
1&-\omega &-i&i\omega &-1&\omega &i&-i\omega \\
1&-1&1&-1&1&-1&1&-1\\
1&-i\omega &i&\omega &-1&i\omega &-i&-\omega \\
1&-i&-1&i&1&-i&-1&i\\
1&\omega &-i&-i\omega &-1&-\omega &i&i\omega \\
\end{bmatrix}}}$$

Links: [wikipedia](https://en.wikipedia.org/wiki/DFT_matrix)
