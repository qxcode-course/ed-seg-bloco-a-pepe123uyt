package main

import (
	"fmt"
	//"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	str := "[ "

	for ll := l.Front(); ll != l.End(); ll = ll.Next() {
		str += fmt.Sprintf("%d", ll.Value)
		if ll.Value == sword.Value {
			str += ">"
		}
		str += " "
	}

	return str + "]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it.Next() == l.End() {
		it = l.Front()
	} else {
		it = it.Next()
	}
	return it
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)

	l := NewDList[int]()

	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}

	sword := l.Front()

	for range chosen - 1 {
		sword = Next(l, sword)
	}

	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}

	fmt.Println(ToStr(l, sword))
}
