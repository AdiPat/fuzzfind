# FuzzFind 🕵️‍♂️

A Fuzzy Search library written in Go.

## Overview 📖

FuzzFind is a powerful and efficient fuzzy search library implemented in Go. It allows you to perform approximate string matching, making it easier to find relevant results even when the search terms are not exact.

## What is Fuzzy Search? 🔍

Fuzzy search is a technique used in information retrieval to find strings that are approximately equal to a given pattern. It is particularly useful when dealing with typos, misspellings, or variations in data.

## Setup 🛠️

To install FuzzFind, use the following command:

```sh
go get github.com/AdiPat/fuzzfind
```

## Usage 🚀

Here's a basic example of how to use FuzzFind:

```go
package main

import (
    "fmt"
    "github.com/yourusername/fuzzfind"
)

func main() {
    searcher := fuzzfind.NewSearcher([]string{"apple", "banana", "grape", "orange"})
    results := searcher.Search("appl")
    fmt.Println(results) // Output: ["apple"]
}
```

## API 📚

### `NewSearcher(data []string) *Searcher`

Creates a new searcher instance with the provided data.

### `(*Searcher) Search(query string) []string`

Performs a fuzzy search on the data and returns a list of matching results.

## Authors 🖋️

- Aditya Patange - [@AdiPat](https://github.com/AdiPat)

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
