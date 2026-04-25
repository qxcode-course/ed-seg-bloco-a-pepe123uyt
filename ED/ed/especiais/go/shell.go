package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	cont := make(map[int]int)

	for _, v := range vet {
		if v < 0 {
			v = -v
		}
		cont[v]++
	}

	var res []Pair
	
	for k, v := range cont {
		res = append(res, Pair{k, v})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].One < res[j].One
	})

	return res
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}

	var res []Pair

	atual := vet[0]
	if atual < 0 {
		atual = -atual
	}
	
	count := 1

	for i := 1; i < len(vet); i++ {
		v := vet[i]
		if v < 0 {
			v = -v
		}

		if v == atual {
			count++
		} else {
			res = append(res, Pair{atual, count})
			atual = v
			count = 1
		}
	}

	res = append(res, Pair{atual, count})

	return res
}

func mnext(vet []int) []int {
	var men []int

	for i := 0; i < len(vet); i++ {

		esq := i > 0 && vet[i-1] < 0
		dir := i < len(vet)-1 && vet[i+1] < 0
		
		if vet[i] > 0 &&  (esq || dir) {
			men = append(men, 1)
		} else {
			men = append(men, 0)
		}

	}
	return men
}

func alone(vet []int) []int {
	var alone []int

	for i := 0; i < len(vet); i++ {
		if i > 0 && vet[i-1] < 0 || i < len(vet)-1 && vet[i+1] < 0 || vet[i] <= 0{
			alone = append(alone, 0)
		} else {
			alone = append(alone, 1)
		}
	}
	return alone
}

func couple(vet []int) int {
	freq := make(map[int]int)

	for _, v := range vet {
		freq[v]++
	}

	casais := 0

	for v, qtd := range freq {
		if v > 0 {
			oposto := freq[-v]

			if qtd < oposto {
				casais += qtd
			} else {
				casais += oposto
			}
		}
	}

	return casais
}


func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos + len(seq) > len(vet){
		return false
	}

	for i := 0; i < len(seq); i++{
		if vet[pos+i] != seq[i] {
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	for i := 0; i <= len(vet)-len(seq); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	for i := 0; i <= len(vet)-len(seq); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {	
	var aux []int
	posMap := make(map[int]bool)

	for _, p := range posList {
		posMap[p] = true
	}

	for i := 0; i < len(vet); i++ {
		if posMap[i]{
			continue
		}

		aux = append(aux, vet[i])
	}
	return aux
}

func clear(vet []int, value int) []int {
	var aux []int

	for _, v := range vet {
		if v != value {
			aux = append(aux, v)
		}
	}
	return aux
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
