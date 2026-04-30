package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type MultiSet struct {
	data 		[]int
	size 		int
	capacity 	int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:	make([]int, capacity),
		size: 	0,
		capacity: capacity,
	}
}

func (m *MultiSet) expand(){
	if m.capacity == 0 {
		m.capacity = 1
	} else {
		m.capacity *= 2
	}

	newData := make([]int, m.capacity)
	copy(newData, m.data)
	m.data = newData

}
func (m *MultiSet) Insert(value int) {
	if m.size == m.capacity {
		m.expand()
	}

	index := 0
	for index < m.size && m.data[index] < value {
		index++
	}

	m.insert(value, index)
}

func (m *MultiSet) insert(value int, index int) error {
	if index < 0 || index > m.size {
		return fmt.Errorf("index out of range")
	}

	for i := m.size; i > index; i-- {
		m.data[i] = m.data[i-1]
	}

	m.data[index] = value
	m.size++
	return nil
} 

func (m *MultiSet) String() string {
	str := "["

	for i := 0; i < m.size; i++ {
		str += fmt.Sprintf("%d", m.data[i])

		if i != m.size-1 {
			str += ", "
		}
	}

	str += "]"
	return str
}

func (m *MultiSet) Contains(value int) bool {
	for _, ms := range m.data {
		if ms == value {
			return true
		}
	}
	return false
}

func (m *MultiSet) Erase(value int) error {
	if !m.Contains(value){
		return fmt.Errorf("value not found")
	}

	index := 0
	for index < m.size && m.data[index] < value {
		index++
	}

	m.erase(index)

	return nil
}

func (m *MultiSet) erase(index int) error {
	if index < 0 || index >= m.size {
		return fmt.Errorf("index out of range")
	}
	for i := index; i < m.size - 1; i++ {
    	m.data[i] = m.data[i+1]
	}

	m.size--
	return nil
}

func (m *MultiSet) Count(value int) int {
	count := 0

	for _, v := range m.data {
		if v == value {
			count++
		}
	}

	return count
}

func (m *MultiSet) Unique() int {
	if m.size == 0 {
		return 0
	}

	count := 1

	for i := 1; i < m.size; i++ {
		if m.data[i] != m.data[i-1] {
			count++
		}
	}

	return count
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (m *MultiSet) Clear() {
	m.data = nil
	m.size = 0
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, err := strconv.Atoi(args[1])
			if err != nil {
    			fmt.Println("invalid number")
    			return
			}

			if err := ms.Erase(value); err != nil {
    			fmt.Println(err)
			}	

		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			count := ms.Count(value)
			fmt.Println(count)
		case "unique":
			count := ms.Unique()
			fmt.Println(count)
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
