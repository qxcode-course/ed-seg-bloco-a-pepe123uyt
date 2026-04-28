package main

import "fmt"

func eh_primo(x int, div int) bool {
	if x < 2 {
		return false
	}

	if div * div > x{
		return true
	}
	if x % div == 0 {
		return false
	}
	return eh_primo(x, div+1);
}
func ene_primo(num, base int) int{
    if base == num && eh_primo(base, 2) {
        return base
    }

    return ene_primo(num, base+1)
    
}
func main() {
	var num int
	fmt.Scan(&num)

	fmt.Println(ene_primo(num, 2))
5
}
