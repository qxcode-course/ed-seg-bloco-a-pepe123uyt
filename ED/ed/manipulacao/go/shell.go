package main

import (
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

func mergeSort(vet []int) []int {
	if len(vet) <= 1 {
		return vet
	}
	meio := len(vet)/2

	esq := vet[0:meio]
	dir := vet[meio:]

	esq = mergeSort(esq)
	dir = mergeSort(dir)

	return mergeA(esq, dir)

}

func mergeA(esq, dir []int)[]int {
	var vet []int
	i, j := 0, 0

	for i < len(esq) && j < len(dir) {
		if esq[i] < dir[j] {
			vet = append(vet, esq[i])
			i++
		} else {
			vet = append(vet, dir[j])
			j++
		}
	}

	for i < len(esq){
		vet = append(vet, esq[i])
		i++
	}
	for j < len(dir){
		vet = append(vet, dir[j])
		j++
	}
	return vet
}

func sortVet(vet []int) []int {
	// for i := 0; i < len(vet); i++{
	// 	for j := 0; j < len(vet); j++{
	// 		if vet[i] < vet[j] {
	// 			vet[i], vet[j] = vet[j], vet[i]
	// 		}
	// 	}
	// }
	return mergeSort(vet)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func mergeStres(vet []int) []int{
	if len(vet) <= 1 {
		return vet
	}
	meio := len(vet)/2

	esq := vet[0:meio]
	dir := vet[meio:]

	esq = mergeStres(esq)
	dir = mergeStres(dir)

	return mergeB(esq, dir)
}

func mergeB(esq, dir []int) []int {
	var vet []int
	i, j := 0, 0

	for i < len(esq) && j < len(dir) {
		if abs(esq[i]) <= abs(dir[j]) {
			vet = append(vet, esq[i])
			i++
		} else {
			vet = append(vet, dir[j])
			j++
		}
	}

	for i < len(esq){
		vet = append(vet, esq[i])
		i++
	}
	for j < len(dir){
		vet = append(vet, dir[j])
		j++
	}
	return vet


}
func sortStress(vet []int) []int {
	// for i := 0; i < len(vet); i++{
	// 	for j := 0; j < len(vet) - 1; j++{
	// 		if abs(vet[j]) > abs(vet[j+1]) {
    //             vet[j], vet[j+1] = vet[j+1], vet[j]
    //         }
	// 	}
	// }
    return mergeStres(vet)
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

	vt := make(map[int]bool)
	var rep []int

	for _, v := range vet{
		_, existe := vt[v]

		if !existe{
			vt[v] = true
		} else {
			rep = append(rep, v)
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

