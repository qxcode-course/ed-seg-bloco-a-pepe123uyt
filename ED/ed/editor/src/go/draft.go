package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Value rune
	next  *Node
	prev  *Node
	root  *Node
}

// Retorna o próximo nó
func (n *Node) Next() *Node {
	return n.next
}

// Retorna o nó anterior
func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	root *Node
	size int
}

func NewList() *List {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root
	return &List{root: root, size: 0}
}

func (l *List) PushBack(value rune) {
	l.Insert(l.root, value)
}

func (l *List) Insert(it *Node, value rune) *Node {
	n := &Node{Value: value, root: l.root}
	n.prev = it.prev
	n.next = it
	it.prev.next = n
	it.prev = n
	l.size++
	return n
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *List) End() *Node {
	return l.root
}

func (l *List) Front() *Node {
	if l.size == 0 {
		return l.root
	}
	return l.root.next
}

func (l *List) Back() *Node {
	if l.size == 0 {
		return l.root
	}
	return l.root.prev
}

func (l *List) Erase(n *Node) *Node {
	if n == l.root {
		return l.root
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	l.size--
	return n.next
}

func (l *List) IndexOf(n *Node) int {
	if n == l.root {
		return l.size
	}
	if n.root != l.root {
		return -1
	}
	i := 0
	for it := l.root.next; it != n; it = it.next {
		i++
		if i >= l.size {
			return -1
		}
	}
	return i
}

func (l *List) String() string {
	values := []string{}
	for n := l.root.next; n != l.root; n = n.next {
		values = append(values, fmt.Sprint(n.Value))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
}


func main(){

	fmt.Println()
}