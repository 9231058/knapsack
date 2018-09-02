# knapsack
[![Travis branch](https://img.shields.io/travis/1995parham/knapsack/master.svg?style=flat-square)](https://travis-ci.org/1995parham/knapsack)

## Intorduction
The knapsack problem is an integer programming problem with boolean variables.
We relax its binary variables with real numbers from 0 to 1 and solve
the resulted non-convex problem by SCA methods.

Implementation has done in Go beacuase of its awesome performance
and Cplex beacuse of its awesome features.

## Cplex
After relaxing binary variables and add following constraint:

`x^2 - x <= 0`

We have non-convex problem, for solving this problem we use lagrangian method create
following problem:

```latex
max sum_i=0^T x_i + mu * (x_i^2 - x_i)
sum_i=0^T x_i <= capacity
```

Again we approximate `x^2 - x` with its first order approximation and after that we
have following LP problem. please note that `x^n_i` means `x_i` in `n`th iteration.

```latex
max sum_i=0^T x^n_i + mu * (2 * x^(n-1)_i - 1) * x^n_i
sum_i=0^T x^n_i <= capacity
```

Implementation of the above LP Problem can be found in `/cplex`.

## Contributers
- Parham Alvani (MSc Student of the Amirkabir University of Technology)
- Bahador Bakhshi (Assistant Professor of the Amirkabir University of Technology)
