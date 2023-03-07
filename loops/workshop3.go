package loops

import "fmt"

func Demo5() {
	var num1 int = 0
	var num2 int = 0
	var result1 int = 0
	var result2 int = 0
	fmt.Println("Input number 1:")
	fmt.Scan(&num1)
	fmt.Println("Input number 2:")
	fmt.Scan(&num2)
	for i := 1; i < num1; i++ {
		if num1%i == 0 {
			result1 += i
		}
	}
	for i := 1; i < num2; i++ {
		if num2%i == 0 {
			result2 += i
		}
	}
	if result1 == num2 && result2 == num1 {
		fmt.Printf("%v and %v are friend numbers.", num1, num2)
	} else {
		fmt.Printf("%v and %v are not friend numbers.", num1, num2)
	}
}
