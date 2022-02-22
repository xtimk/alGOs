[![Build Status](https://app.travis-ci.com/xtimk/alGOs.svg?token=9yNzkmTjR26qHy4qyCXB&branch=master)](https://app.travis-ci.com/xtimk/alGOs) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/xtimk/alGOs/blob/master/LICENSE)
# GoLang Implementation of classic & advanced algorithms
Classic and advanced algorithms implemented in Go.

## Building
```go
go build
```
This generates an executable file named `alGOs` (which is the main loop from the `algos.go` file)

## Testing
```bash
./run_tests.sh
```

## Run
```bash
./alGOs
```

## Algorithms
Here are listed the implemented algorithms

|Algorithm|Complexity parameters|Worst-Case time complexity|Avg-case time complexity|Best-case time complexity|Auxiliary space complexity|
|---------|---------------------|-------------------------|------------------------|-|-|
|Naive Exact Pattern-Matching|`t=\|Text\|` <br> `p=\|Pattern\|`|`Θ(t*p)`|`O(t*p)`|`O(1)`|
