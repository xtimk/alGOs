[![Build Status](https://app.travis-ci.com/xtimk/alGOs.svg?token=9yNzkmTjR26qHy4qyCXB&branch=master)](https://app.travis-ci.com/xtimk/alGOs) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/xtimk/alGOs/blob/master/LICENSE)
# GoLang Implementation of classic & advanced algorithms
Classic and advanced algorithms implemented in GoLang.

## Testing the package
```bash
go test github.com/xtimk/alGOs/algos
```

## Run the main example
```bash
go run main_example.go
```

## Algorithms
Here are listed the implemented algorithms

|Algorithm|Description|Complexity parameters|Worst-Case time complexity|Avg-case time complexity|Best-case time complexity|Auxiliary space complexity|
|---------|-|--------------------|-------------------------|------------------------|-|-|
|Naive Exact Pattern-Matching|Pattern matching algorithm|`t=\|Text\|`  <br>  `p=\|Pattern\|`|`Θ(t*p)`|`O(t*p)`|`O(t)`|`O(1)`|
|LPS|Longest Prefix that is also a suffix|`t=\|Text\|`|`Θ(t)`|`O(t)`|`O(t)`|`O(1)`|
|KMP|Knuth Morris Pratt algorithm for exact pattern matching. Makes use of the LPS algorithm up here|`t=\|Text\|`  <br>  `p=\|Pattern\|`|`Θ(t+p)`|`O(t+p)`|`O(t)`|`O(p)`|
