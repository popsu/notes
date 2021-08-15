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

To convert back use the inverse:

$${\displaystyle V^{-1} = \frac{\overline V}{n}={ \frac{1}{n} \cdot \begin{bmatrix}
1&1&1&1&1&1&1&1\\
1&i\omega &i&-\omega &-1&-i\omega &-i&\omega \\
1&i&-1&-i&1&i&-1&-i\\
1&-\omega &-i&i\omega &-1&\omega &i&-i\omega \\
1&-1&1&-1&1&-1&1&-1\\
1&-i\omega &i&\omega &-1&i\omega &-i&-\omega \\
1&-i&-1&i&1&-i&-1&i\\
1&\omega &-i&-i\omega &-1&-\omega &i&i\omega \\
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

QED

Links: [wikipedia](https://en.wikipedia.org/wiki/DFT_matrix)
