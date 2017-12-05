# troxel-go
Troxel GO Library

## Download

```
go get github.com/yacklebeam/troxel-go
```

## Import

```
import (
    "github.com/yacklebeam/troxel-go"
)
```

## Functions

Default namespace is "trx".

> **ReadFlatFileToSlice (filename string) []string**

Reads *filename* flat file and puts each line into the slice.

> **ToSet (list []string) []string**

Returns set (slice) of values from *list*

> **SetDifference (X, Y []string) []string**

Returns difference of sets X and Y (X - Y)

> **SetUnion (X, Y []string) []string**

Returns union of sets X and Y (X v Y)

> **SetIntersection (X, Y []string) []string**

Returns intersection of sets X and Y (X ^ Y)

## Example

```
package main

import (
    "github.com/yacklebeam/troxel-go"
    "fmt"
)

func main() {
    var X, Y []string
    X = trx.ToSet(trx.ReadFlatFileToSlice("test.txt"))
    Y = trx.ToSet(trx.ReadFlatFileToSlice("test2.txt"))
    
    fmt.Println(trx.SetDifference(X, Y))
    fmt.Println(trx.SetDifference(Y, X))
    fmt.Println(trx.SetUnion(X, Y))
    fmt.Println(trx.SetIntersection(X, Y))
}
```