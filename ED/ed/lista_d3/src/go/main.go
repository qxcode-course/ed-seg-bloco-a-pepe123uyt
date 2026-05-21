package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n

	l.size++
}

func (l *LList) String(ll *LList) string {
	str := "["

	for n := ll.root.next; n != ll.root; n = n.next {
		str += fmt.Sprintf("%d", n.Value)
		if n != ll.root.prev {
			str += ", "
		}
	}
	str += "]"
	
	return str
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}

func equals(lla, llb *LList) bool {
	if lla.size != llb.size {
		return false
	}

	for i, j := lla.root.next, llb.root.next; i != lla.root; i, j = i.next, j.next {
		if i.Value != j.Value {
			return false
		}
	}

	return true
	
}

func addsorted(lla *LList, value int) {
	n := lla.root.next

	for n != lla.root && n.Value < value {
		n = n.next
	}

	lla.insertBefore(n, value)
	lla.size++

}

func reverse(ll *LList) {
	node := ll.root.next

	for node != ll.root {
		next := node.next
		node.next, node.prev = node.prev, node.next
		node = next
	}

	ll.root.next, ll.root.prev = ll.root.prev, ll.root.next
}

func merge(lla, llb *LList) *LList {
	for n := llb.root.next; n != llb.root; n = n.next {
		addsorted(lla, n.Value)
	}

	return lla
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			 lla := str2list(args[1])
			 llb := str2list(args[2])
			 if equals(lla, llb) {
			 	fmt.Println("iguais")
			 } else {
			 	fmt.Println("diferentes")
			 }
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
			 	value, _ := strconv.Atoi(args[i])
			 	addsorted(lla, value)
			}
			fmt.Println(lla.String(lla))
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla.String(lla))
		case "merge":
		 	lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged.String(merged))
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
