package main

import "fmt"

func main() {
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
	fmt.Println(triple + penta + fifty)
}
