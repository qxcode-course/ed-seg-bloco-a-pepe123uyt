package main

import (
	"fmt"
)

func ganhou(times *Queue[string], mm, nn int) {
	t1 := times.Dequeue()
	t2 := times.Dequeue()

	if mm > nn {
    	times.Enqueue(t1)
	} else {
		times.Enqueue(t2)
	}
}

func main() {
	times := NewQueue[string]()
	for i := 65; i <= 80; i++{
		times.Enqueue(string(i))
	}
	
	for range 15 {
		var mm, nn int
		fmt.Scan(&mm, &nn)
		ganhou(times, mm, nn)
	}
	
	fmt.Printf("%s\n", times.items.Front().Value)
	// times, resul = oitavas(times, resul)
	// times, resul = quartas(times, resul)
	// times, resul = semi(times, resul)
	// times, resul = final(times, resul)


}
