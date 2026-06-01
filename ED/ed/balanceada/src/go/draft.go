package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
    data []rune
}
func NewStack() *Stack{
    return &Stack{data: []rune{}}
}
func (s *Stack) Push(value rune) {
    s.data = append(s.data, value)
}
func (s *Stack) Top() rune {
	return s.data[len(s.data)-1]
}
func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}
func (s *Stack) IsEmpty() bool {
    if len(s.data) <= 0 {
        return true
    }
    return false
}

func isBalance(formula []rune) bool {
    if len(formula) % 2 != 0 {
        return false
    }

    stack := NewStack()
    
    for _, i := range formula {
        if i == '(' || i == '[' {
            stack.Push(i)
        }

        if !stack.IsEmpty() {
            s := stack.Top()
    
            if i == ')'{
                if s == '(' || s == ']'{
                    stack.Pop()
                } else {
                    return false
                }
            }
            if i == ']' {
                if s == '[' || s == ')' {
                    stack.Pop()
                } else {
                    return false
                }
            }
            
        } else {
            return false
        }

    }
    return true
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    formula := []rune(scanner.Text())

    if isBalance(formula) {
        fmt.Println("balanceado")
    } else {
        fmt.Println("nao balanceado")
    }
}