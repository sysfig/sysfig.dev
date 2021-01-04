Memory
  - heap - global, dynamically allocated
  - stack - local to the scope of the function

Go manages memory allocation automatically. There's no mechanism for manual memory allocation.

Stack allocation is cheap because it requires only two CPU instructions:
  - push onto the stack for allocation
  - release from the stack

Because the variable ??, the lifetime and memory requirements of variable on the stack can be determined at compile time??. On the other hand, variables allocated on the stack 

Heap allocation is expensive

https://golang.org/ref/mem

## Concurrency

## Data Race Detector

https://golang.org/doc/articles/race_detector.html
