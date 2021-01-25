package main

import "fmt"

func main() {

	fmt.Println("problem1 =", problem1())
	fmt.Println("problem2 =", problem2())
}

func problem1() int {
	triple := 0
	penta := 0
	fifty := 0
	for i := 1; i < 1000; i++ {
		if i%15 == 0 {
			fifty += i
		} else if i%3 == 0 {
			triple += i
		} else if i%5 == 0 {
			fifty += i
		}
	}
	return triple + penta + fifty
}

func problem2() int {
	var memo [2000]int
	memo[0] = 1
	memo[1] = 1
	sum := 0
	for i := 2; i < 2000; i++ {
		memo[i] = memo[i-2] + memo[i-1]
		if memo[i] > 4000000 {
			break
		}
		if memo[i]%2 == 0 {
			sum += memo[i]
		}
	}
	return sum
}
