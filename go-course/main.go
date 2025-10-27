package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var investmentAmount = 1000.0
	var expectedReturn = 0.05 // 5% expected return
	var years = 10

	var futureValue = investmentAmount * (1 + expectedReturn)
}
