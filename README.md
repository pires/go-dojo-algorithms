# go-dojo-algorithms
Go experiments with algorithms.

## [Grokking Algorithms](https://www.manning.com/books/grokking-algorithms)

This book is an amazing resource for a newbie like me.

Here's the source code for each chapter:

* [Chapter 1](./grokking_algorithms/ch01)
  * Binary Search
* [Chapter 2](./grokking_algorithms/ch02)
  * Selection Sort

### Tests

```
$ go test -cover ./grokking_algorithms/...
ok  	github.com/pires/go-dojo-algorithms/grokking_algorithms/ch01	0.014s	coverage: 95.8% of statements
ok  	github.com/pires/go-dojo-algorithms/grokking_algorithms/ch02	0.009s	coverage: 100.0% of statements
ok  	github.com/pires/go-dojo-algorithms/grokking_algorithms/util	0.012s	coverage: 80.0% of statements
```
