package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func sum(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total
}

func difference(numbers ...float64) float64 {
	total := 0.0
	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			total = numbers[i]
		} else {
			total -= numbers[i]
		}
	}
	return total
}

func product(numbers ...float64) float64 {
	total := 1.0
	for _, number := range numbers {
		total *= number
	}
	return total
}

func main() {
	/*
		var f float64
		fmt.Println("This is a simple command line calculator. Please enter the expression you would like to evaluate and hit ENTER.")
		//reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter numbers to add: ")
		_, err := fmt.Scanf("%f", &f)
		if err != nil {
			fmt.Println("You encountered an error:", err)
		} else {
			fmt.Println(sum(f))
		}
		//expression, _ := reader.ReadString('\n')
		//fmt.Println(sum(inputs))
	*/

	// ********

	//Learning how to use GoConvey with simple functions.
	fmt.Println(sum(1, 2, 3))
	fmt.Println(difference(1, 2, 3))
	fmt.Println(product(1, 2, 3))
}
