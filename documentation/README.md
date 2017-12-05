# Documentation

Default namespace is "trx".

### File Reading

Function | Parameters | Return | About
--- | --- | --- | ---
ReadFlatFileToSlice | filename *string* | []string | Reads *filename* flat file and puts each line into the slice.
ReadCSVFileToList | filename *string* | [][]string | Reads CSV file, and returns a list of lists of the items.

### Set Functions

Function | Parameters | Return | About
--- | --- | --- | ---
ToSet | list *[]string* | []string | Returns set (slice) of values from *list*
SetDifference | X, Y *[]string* | []string | Returns difference of sets X and Y (X - Y)
SetUnion | X, Y *[]string* | []string | Returns union of sets X and Y (X v Y)
SetIntersection | X, Y *[]string* | []string | Returns intersection of sets X and Y (X ^ Y)
