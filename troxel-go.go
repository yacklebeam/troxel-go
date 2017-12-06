package trx

import (
    "bufio"
    "log"
    "os"
    "strings"
)

func FlattenStr(arr [][]string) []string {
    var ret []string

    for _,x := range arr {
        for _,y := range x {
            ret = append(ret, y)
        }
    }

    return ret
}

func FlattenInt(arr [][]int) []int {
    var ret []int

    for _,x := range arr {
        for _,y := range x {
            ret = append(ret, y)
        }
    }

    return ret
}

func ContainsStr(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func ContainsInt(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func ReadFlatFile(filename string) []string {
    file, err := os.Open(filename)

    var ret []string

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var text = scanner.Text()
        ret = append(ret, text)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return ret
}

func ToSet(list []string) []string {
    var ret []string

    for _, x := range list {
        if !ContainsString(ret, x) {
            ret = append(ret, x)
        }
    }

    return ret
}

func SetDifference(X, Y []string) []string {
    var ret []string

    for _, x := range X {
        if !ContainsString(Y, x) {//in x, but not y
            ret = append(ret, x)
        }
    }

    return ret
}

func SetUnion(X, Y []string) []string {
    var ret []string

    for _, x := range X {
        if !ContainsString(ret, x) {
            ret = append(ret, x)
        }
    }

    for _, y := range Y {
        if !ContainsString(ret, y) {
            ret = append(ret, y)
        }
    }

    return ret
}

func SetIntersection(X, Y []string) []string {
    var ret []string

    for _, x := range X {
        if ContainsString(Y, x) {//in x, and in y
            ret = append(ret, x)
        }
    }

    return ret
}

func ReadCSVFile(filename string) [][]string {
    file, err := os.Open(filename)

    var ret [][]string

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var text = scanner.Text()
        var item = strings.Split(text, ",")
        ret = append(ret, item)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return ret
}