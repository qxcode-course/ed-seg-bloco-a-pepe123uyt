package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}
func (v *Set) String() string {
	str := "["

	for i := 0; i < v.size; i++ {
		if i == v.size - 1{
			str += fmt.Sprintf("%d", v.data[i])
			break
		}
		str += fmt.Sprintf("%d, ",v.data[i])
	}
	str += "]"

	return str
}
func (v *Set) Insert(value int) {
	index := v.size

	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return 
		}
		if v.data[i] > value {
			index = i
			break
		}
	} 

	if v.capacity == 0 {
		v.capacity = 1
		v.data = make([]int, v.capacity)
	} else if v.size == v.capacity {
		v.capacity *= 2
		newData := make([]int, v.capacity)
		copy(newData, v.data)
		v.data = newData
	}
	
	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	
	v.data[index] = value
	v.size++
}

func (v *Set) Contains(value int) bool {
	for _, c := range v.data {
		if c == value {
			return true
		}
	}
	return false
}


func (v *Set) Erase(value int) bool {
	index := -1

	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}

	v.size--
	return true
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			str := v.String()
			fmt.Println(str)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value){
				v.Erase(value)
			} else {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value){
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
