package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Size() int {
	return v.size
}
func (v *Vector) Capacity() int {
	return v.capacity
} 

func (v *Vector) Status() string {
	var str string

	str += fmt.Sprintf("size:%d capacity:%d", v.Size(), v.Capacity())

	return str
}

func (v *Vector) String() string {
	str := "["

	for i := 0; i < v.Size(); i++ {
		if i == v.Size() - 1{
			str += fmt.Sprintf("%d", v.data[i])
			break
		}
		str += fmt.Sprintf("%d, ",v.data[i])
	}
	str += "]"
	return str
} 
func (v *Vector) Reserve(newCapacity int) {
	if newCapacity <= v.capacity {
		return
	}

	newData := make([]int, newCapacity)

	for i := 0; i < v.size; i++ {
		newData[i] = v.data[i]
	}

	v.data = newData
	v.capacity = newCapacity
}

func (v *Vector) PushBack(value int) {
    if v.Size() == v.Capacity() {

		newCapacity := 1
		if v.Capacity() > 0 {
			newCapacity = v.Capacity() * 2
		}

		newData := make([]int, newCapacity)

		for i := 0; i < v.Size(); i++ {
			newData[i] = v.data[i]
		}

		v.data = newData
		v.capacity = newCapacity
	}

	v.data[v.size] = value
	v.size++
}

func (v *Vector) Get(index int) int {
	return v.data[index]
}

func (v *Vector) At(index int) (int, error) {
	if index < 0 || index >= v.size {
		return 0, fmt.Errorf("index out of range")
	}
	return v.data[index], nil
}

func (v *Vector) Set(index int, value int) error{
	if index > v.Size() - 1 {
		return fmt.Errorf("index out of range")
	}
	v.data[index] = value
	return nil
}

func (v *Vector) Clear() {
	v.size = 0
	v.data = nil
}

func (v *Vector) PopBack() (int, error) {
	if v.Size() == 0 {
		return 0, fmt.Errorf("vector is empty")
	}	
	last := v.data[v.Size()-1]

	v.size--

	return last, nil
}

func (v *Vector) Insert(index int, value int) error {
	if index < 0 || index > v.Size() {
		return fmt.Errorf("index out of range")
	}

	if v.Size() == v.Capacity() {
		v.capacity *= 2
		newData := make([]int, v.capacity)
		copy(newData, v.data)
		v.data = newData
	}

	for i := v.Size(); i > index; i-- {
		v.data[i] = v.data[i-1]
	}

	v.data[index] = value
	v.size++

	return nil
}

func (v *Vector) Erase(index int) error {
	if index < 0 || index >= v.Size() {
		return fmt.Errorf("index out of range")
	}

	for i := index; i < v.Size()-1; i++ {
    v.data[i] = v.data[i+1]
	}

	v.size--
	return nil
}

func (vet *Vector) Contains(value int) bool{
	for _, v := range vet.data {
		if v == value {
			return true
		}
	}
	return false
} 

func (v *Vector) IndexOf(value int) int {
	if !v.Contains(value){
		return -1
	}

	var index, i int

	for _, valor := range v.data {
		if valor == value {
			index = i
			break
		}
		i++
	}
	return index
	
}   

func (v *Vector) Slice(start int, end int) *Vector {
	if v.size == 0 {
		return NewVector(0)
	}

	start = ((start % v.size) + v.size) % v.size
	end = ((end % v.size) + v.size) % v.size

	if start <= end {
		return &Vector{
			data:     v.data[start:end],
			size:     end - start,
			capacity: end - start,
		}
	}

	return nil

} 
func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v.String())
		case "status":
			fmt.Println(v.Status())
		case "pop":
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}
			
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
