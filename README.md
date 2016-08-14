# go-kCombinations
Memory efficient k-combinations functions for golang primitives

## Example

```` go
for combo := range kCombinations.GenerateCombinationsString([]string{"1","2","3","4"}, 2) {
    fmt.Println(combo)
}
````

## Badges

[![GoDoc](https://godoc.org/github.com/Ramshackle-Jamathon/go-kCombinations?status.svg)](https://godoc.org/github.com/Ramshackle-Jamathon/go-kCombinations)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)
