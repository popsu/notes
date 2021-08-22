# Abstract Algebra

## Field

Field is a set $F$ together with two binary operations _addition_ ($+$) and _multiplication_ ($\cdot$) that satisfy the following properties:

1. **Associativity** of addition and multiplication: $a + (b + c) = (a + b) + c$, and $a \cdot (b \cdot c) = (a \cdot b) \cdot c$.
2. **Commutativity** of addition and multiplication: $a + b = b + a$, and $a \cdot b = b \cdot a$
3. Additive and multiplicative **identity**: there exists different elements $0$ and $1$ such that $a + 0 = a$ and $a \cdot 1 = a$
4. **Additive inverses**: for every $a$ in $F$ there exists an element in $F$: $-a$ such that $a + (-a) = 0$.
5. **Multiplicative inverses**: for every $a \neq 0$ in $F$ there exists an element in $F$: $a^{-1}$ such that $a \cdot a^{-1} = 1$.
6. **Distributivity** of multiplication over addition: $a \cdot (b + c) = (a \cdot b) + (a \cdot c)$.

Summarized a field is a commutative ring where $0 \neq 1$ and all nonzero elements are invertible.

## Finite field

A field with finite number of elements is called finite field (or Galois field). The number of elements is always of form $p^k$, where $p$ is a prime number, and $k$ is a positive integer. For example there is **no** finite field with $10 = 2 \cdot 5$ elements. But there **is** one with $256 = 2 ^ 8$ elements.

Uniqueness of finite fields: all finite fields with the same number of elements are isomorphic.

## Mutliplicative group

In a field, the non-zero elements with multiplication form a group. This is called multiplicative group.

## Fundamental theorem of finite abelian groups

Let $G$ be a finite abelian group. Then $G$ is a direct product of cyclic groups of prime-power order.

??? example "Example - Number of different abelian groups with 1008 elements?"
    Prime factorization: $1008 = 2^4 \cdot 3^2 \cdot 7$. Looking at the powers we get 4, 2, 1. We get the result by multiplying the different [partitions](https://en.wikipedia.org/wiki/Partition_(number_theory)) of these numbers.

    4 can be partitioned 5 ways: $4, 3+1, 2+2, 2+1+1, 1+1+1+1$

    2 can be partitioned 2 ways: $2, 1+1$

    1 can be partitioned 1 way: $1$

    So overall $5 \cdot 2 \cdot 1 = 10$ different abelian groups with 1008 elements.

??? example "Example - Number of different abelian groups with 256 elements?"
    Prime factorization: $256 = 2^8$. There are 22 partitions of 8, so there are 22 different abelian groups with 256 elements.

## Links

[Finite field - wikipedia](https://en.wikipedia.org/wiki/Finite_field)
[301.11D The Fundamental Theorem of Finite Abelian Groups - youtube](https://www.youtube.com/watch?v=LPDlaG3fAfQ&list=PLL0ATV5XYF8AQZuEYPnVwpiFy0jEipqN-)
