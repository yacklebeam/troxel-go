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
    
    X = trx.ToSet(trx.ReadFlatFile("test.txt"))
    Y = trx.ToSet(trx.ReadFlatFile("test2.txt"))
    
    //Using Set Functions
    fmt.Println(trx.SetDifference(X, Y))
    fmt.Println(trx.SetDifference(Y, X))
    fmt.Println(trx.SetUnion(X, Y))
    fmt.Println(trx.SetIntersection(X, Y))

    //Reading CSV File
    var Z [][]string
    Z = trx.ReadCSVFile("test.csv")
    
    fmt.Println(Z)
}
```