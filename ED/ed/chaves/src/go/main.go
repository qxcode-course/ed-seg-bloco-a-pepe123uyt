package main

import (
	"fmt"
)

func ganhou(times *Queue[string], mm, nn int) string{
	t1 := times.Dequeue()
	t2 := times.Dequeue()

	if mm > nn {
    	vencedor := t1
    	times.Enqueue(vencedor)
    	return vencedor
	}

	vencedor := t2
	times.Enqueue(vencedor)
	return vencedor
}

func main() {
	times := NewQueue[string]()
	for i := 65; i <= 80; i++{
		times.Enqueue(string(i))
	}
	
	for range 15 {
		var mm, nn int
		fmt.Scan(&mm, &nn)
		times.Dequeue()
		times.Enqueue(ganhou(times, mm, nn))
	}
	
	fmt.Println(times)

	// times, resul = oitavas(times, resul)
	// times, resul = quartas(times, resul)
	// times, resul = semi(times, resul)
	// times, resul = final(times, resul)


}
