# anydiff

A Go package that calculates the difference between two slices of any type, including slices of different types.

## Features

- **Generic Support**: Works with slices of any type using Go generics.
- **Custom Comparison**: Accepts user-defined equality functions for flexible comparisons.
- **Edit Script Output**: Returns a minimal sequence of Copy, Addition, and Deletion operations.
- **Efficient Algorithm**: Uses Myers’ O(ND) algorithm for optimal diff computation.

## Installation

```cmd
go get github.com/makiuchi-d/anydiff
```

## Usage

```go
package main

import (
	"fmt"
	"strings"

	"github.com/makiuchi-d/anydiff"
)

type User1 struct {
	ID   int
	Name string
}

type User2 struct {
	ID   int
	Name string
	Age  int
}

func main() {
	a := []User1{
		{1, "alice"},
		{2, "bob"},
		{3, "carol"},
		{4, "dave"},
	}
	b := []User2{
		{1, "ALICE", 30},
		{3, "CHARLIE", 24},
		{4, "DAVE", 41},
		{5, "ELEN", 18},
	}

	// custom comparison: case insensitive, ignore age
	cmp := func(a *User1, b *User2) bool {
		return a.ID == b.ID && strings.ToLower(a.Name) == strings.ToLower(b.Name)
	}

	// compute minimal edit operations
	edit := anydiff.Diff(a, b, cmp)

	// print diff
	i, j := 0, 0
	for _, o := range edit {
		switch o {
		case anydiff.Deletion:
			fmt.Println("-", a[i])
			i++
		case anydiff.Addition:
			fmt.Println("+", b[j])
			j++
		case anydiff.Keep:
			fmt.Println(" ", b[j])
			i++
			j++
		}
	}

	// output:
	//   1: ALICE (30)
	// - 2: bob
	// - 3: carol
	// + 3: CHARLIE (24)
	//   4: DAVE (41)
	// + 5: ELEN (18)
}

func (u User1) String() string {
	return fmt.Sprintf("%d: %s", u.ID, u.Name)
}

func (u User2) String() string {
	return fmt.Sprintf("%d: %s (%d)", u.ID, u.Name, u.Age)
}
```

Run on Playground: [https://go.dev/play/p/kbk3UxwK3Yo](https://go.dev/play/p/kbk3UxwK3Yo)
