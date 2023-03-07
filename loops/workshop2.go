package loops

import "fmt"

func Demo4() {
	var input int = 0
	fmt.Println("Input a number:")
	fmt.Scan(&input)
	isPrime := true
	for i := 2; i < input; i++ {
		if input%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Printf("Number is prime: %v", isPrime)
}
