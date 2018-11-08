package trx

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// FlattenStr takes an array of string arrays, and returns a single array of strings
func FlattenStr(arr [][]string) []string {
	var ret []string

	for _, x := range arr {
		for _, y := range x {
			ret = append(ret, y)
		}
	}

	return ret
}

// FlattenInt takes an array of int arrays, and returns a single array of ints
func FlattenInt(arr [][]int) []int {
	var ret []int

	for _, x := range arr {
		for _, y := range x {
			ret = append(ret, y)
		}
	}

	return ret
}

// ContainsStr takes an array of string s, and a string e and returns True if e is in s, or False otherwise
func ContainsStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsInt64 takes an array of int64 s, and an int64 e and returns True if e is in s, or False otherwise
func ContainsInt64(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsInt takes an array of int s, and an int e and returns True if e is in s, or False otherwise
func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ReadFlatFile take a filename string, and returns an array of strings -- each line from the file is an element in the array, in order.
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

// ToSet takes an array of strings "list", and returns an array of strings of the unqiue values from list
func ToSet(list []string) []string {
	var ret []string

	for _, x := range list {
		if !ContainsStr(ret, x) {
			ret = append(ret, x)
		}
	}

	return ret
}

// SetDifference takes two string arrays X and Y, and returns the set difference of them as a string array
func SetDifference(X, Y []string) []string {
	var ret []string

	for _, x := range X {
		if !ContainsStr(Y, x) { //in x, but not y
			ret = append(ret, x)
		}
	}

	return ret
}

// SetUnion takes two string arrays X and Y, and returns the set union of them as a string array
func SetUnion(X, Y []string) []string {
	var ret []string

	for _, x := range X {
		if !ContainsStr(ret, x) {
			ret = append(ret, x)
		}
	}

	for _, y := range Y {
		if !ContainsStr(ret, y) {
			ret = append(ret, y)
		}
	}

	return ret
}

// SetIntersection takes two string arrays X and Y, and returns the set intersection of them as a string array
func SetIntersection(X, Y []string) []string {
	var ret []string

	for _, x := range X {
		if ContainsStr(Y, x) { //in x, and in y
			ret = append(ret, x)
		}
	}

	return ret
}

// ReadCSVFile takes a filename string, and returns an array of arrays read from a CSV
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
		var realList []string
		tempString := ""
		done := false
		mid := false
		for _, str := range item {
			if len(str) == 0 {
				realList = append(realList, str)
				continue
			}
			if str[0] == '"' && str[len(str)-1] != '"' {
				tempString = tempString + str
				mid = true
			} else if str[0] != '"' && str[len(str)-1] == '"' {
				tempString = tempString + str
				mid = true
				done = true
			} else if !mid {
				tempString = str
				done = true
			} else {
				tempString = tempString + str
			}
			if done {
				mid = false
				done = false
				realList = append(realList, tempString)
				tempString = ""
			}
		}
		ret = append(ret, realList)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ret
}
