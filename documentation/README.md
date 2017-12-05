# Documentation

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

> **ReadCSVFileToList (filename string) [][]string**

Reads CSV file, and returns a list of lists of the items
