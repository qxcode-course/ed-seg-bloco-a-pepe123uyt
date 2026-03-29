package main

import (
	"sort"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	men := []int{}
	for i := 0; i < len(vet); i++{
		if vet[i] > 0 {
			men = append(men, vet[i])
		}
	}

	return men
}

func getCalmWomen(vet []int) []int {
	women := []int{}
	for i := 0; i < len(vet); i++{
		if vet[i] < 0 && vet[i] > -10{
			women = append(women, vet[i])
		}
	}

	return women
}

func sortVet(vet []int) []int {
	novo := make([]int, len(vet))
    copy(novo, vet)
	
    sort.Ints(novo)
    return novo
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func sortStress(vet []int) []int {
	novo := make([]int, len(vet))
    copy(novo, vet)
	
	for i := 0; i < len(novo); i++{
		for j := 0; j < len(novo) - 1; j++{
			if abs(novo[j]) > abs(novo[j+1]) {
                novo[j], novo[j+1] = novo[j+1], novo[j]
            }
		}
	}
    return novo
}

func reverse(vet []int) []int {
	novo := make([]int, len(vet))
	
	for i := 0; i < len(vet); i++{
		novo[i] = vet[len(vet)-1-i]
	}
	return novo
}

func unique(vet []int) []int {
	vt := make(map[int]bool)
	unique := []int{}

	for _, v := range vet{
		if !vt[v]{
			unique = append(unique, v)
			vt[v] = true
		}
	}
	return unique
}

func repeated(vet []int) []int {
	rep := []int{}

	for i := 0; i < len(vet); i++{
			if vet[i-1] == vet[i]{
				rep = append(rep, vet[i])
			}
	}
	return rep
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

