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

**func ReadFlatFileToSlice(filename string) []string**
Reads *filename* flat file and puts each line into the slice.

**func ToSet(list []string) []string**
Returns set (slice) of values from *list*

**func SetDifference(X, Y []string) []string**
Returns difference of sets X and Y (X - Y)

**func SetUnion(X, Y []string) []string**
Returns union of sets X and Y (X v Y)

**func SetIntersection(X, Y []string) []string**
Returns intersection of sets X and Y (X ^ Y)