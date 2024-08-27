## Chapter 3

## Introduction
- *Asymptotic* runnning time referes to the order of growth of running time of an algorithms as the input size increases
- Usually the lower terms of the algorithm's running time is discarded to focus on the rate of growth that is associated with the highest term
- Asymptotic notions are designed to work on functions on general

### Types of asymptotic notations
- Big O notation
  - it denotes the upper bound of an algorithm. It indicates that an algorithm can not grow faster than a certain rate
  - example: the big O bound of the function: 7n^3 + 5n^2 - 12n is O(n^3)
  - suprisingly its big O bound is also: O(n^4)
- Big omega (Ω) notation
  - it denotes the lower bound of an algorithm. It indicates that an algorithm will grow at least as fast as a certain rate
  - example: the big O bound of the function: 7n^3 + 5n^2 - 12n is (Ωn^3)
- Big theta (Θ) notation
  - it denotes the precise running time of the algorithm, that the algorithm will grow exactely as fast at a certain rate
  - if a function is shown to be both O(f(n)) and Ω(f(n)) then that function is automatically is also Θ(f(n))

## Asymptotic notion's formal definations

