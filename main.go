package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("problem1 =", problem1())
	fmt.Println("problem2 =", problem2())
	fmt.Println("problem3 =", problem3())
	fmt.Println("problem4 =", problem4())
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

func problem3() int {
	Number := 600851475143
	i := 2
	max := 0
	for {
		if Number%i == 0 {
			Number /= i
			if max < i {
				max = i
			}
		} else {
			i++
		}
		if Number < i {
			break
		}
	}
	return max
}

func problem4() int{
	max :=0
	for k:=0;k<1000;k++{
		for l := 0;l<1000;l++{
			NumberString := strconv.Itoa(k*l)
			runes := []rune(NumberString)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			ReverseNumberString := string(runes)
			if NumberString == ReverseNumberString{
				if max<k*l{
					max = k*l
				}
			}
		}
	}
	return max
}
