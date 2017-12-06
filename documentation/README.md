# Documentation

Default namespace is "trx".

### File Reading

Function | Parameters | Return | About
--- | --- | --- | ---
ReadFlatFile | filename *string* | []string | Reads flat file *filename* and returns the lines in a list.
ReadCSVFile | filename *string* | [][]string | Reads CSV file *filename* and returns the items in an array.

### Set Functions

Function | Parameters | Return | About
--- | --- | --- | ---
ToSet | list *[]string* | []string | Returns set (no duplicates) of values from *list*
SetDifference | X, Y *[]string* | []string | Returns difference of sets *X* and *Y* (X - Y)
SetUnion | X, Y *[]string* | []string | Returns union of sets *X* and *Y* (X v Y)
SetIntersection | X, Y *[]string* | []string | Returns intersection of sets *X* and *Y* (X ^ Y)
