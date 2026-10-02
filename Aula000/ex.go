package main

import "fmt"

func main() {
	fmt.Println("Olá, mundo!")
	
	var number int
	
	fmt.Scanf("%d", &number)

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}