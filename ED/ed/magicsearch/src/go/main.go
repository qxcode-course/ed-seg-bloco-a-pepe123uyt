package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Exist(slice []int, value int) bool{
	for i := 0; i < len(slice) - 1; i++{
		if slice[i] == value {
			return true
		}
	}
	return false
}
func binarySearch(slice []int, value int) int {
	ini := 0
	fim := len(slice) - 1

	for ini <= fim {
		index := (ini + fim)/2

		if slice[index] == value {
			return index
		}

		if (slice[index] < value) {
			ini = index + 1;
		} else {
			fim = index - 1; 
		}
	}

	return ini
}

func MagicSearch(slice []int, value int) int {
	if len(slice) == 0 {
		return 0
	}

	pos := binarySearch(slice, value)

	if Exist(slice, value) {
		for i := pos; i < len(slice); i++ {
			if i == len(slice) - 1 && slice[i] == value {
				return i
			}
			if slice[i] != value || slice[i] == len(slice) -1{
				return i - 1
			}
		}
	}

	return pos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
