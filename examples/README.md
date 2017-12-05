# Examples

Full Docs: https://github.com/yacklebeam/troxel-go/tree/master/documentation

### File Reading, Set Functions

```
package main

import (
    "github.com/yacklebeam/troxel-go"
    "fmt"
)

func main() {
    
    //Reading Flat File To Sets
    var X, Y []string
    
    X = trx.ToSet(trx.ReadFlatFileToSlice("test.txt"))
    Y = trx.ToSet(trx.ReadFlatFileToSlice("test2.txt"))
    
    //Using Set Functions
    fmt.Println(trx.SetDifference(X, Y))
    fmt.Println(trx.SetDifference(Y, X))
    fmt.Println(trx.SetUnion(X, Y))
    fmt.Println(trx.SetIntersection(X, Y))

    //Reading CSV File
    var Z [][]string
    Z = trx.ReadCSVFileToList("test.csv")
    
    fmt.Println(Z)
}
```