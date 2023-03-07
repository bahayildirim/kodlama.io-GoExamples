package loops

import "fmt"

func Demo3() {
	var num int = 10
	var guessedNum int = 0
	var guessAmount int = 0
	var isGuessed bool = false
	for !isGuessed {
		fmt.Println("Guess a number:")
		fmt.Scanln(&guessedNum)
		if num == guessedNum {
			fmt.Println("You have guessed the number!")
			isGuessed = true
		} else if num <= guessedNum {
			fmt.Println("Number too high")
		} else {
			fmt.Println("Number too low")
		}
		guessAmount++
	}
	fmt.Printf("Amount of guesses: %v ", guessAmount)
	if guessAmount < 4 {
		fmt.Println("Good job!")
	} else if guessAmount > 4 && guessAmount < 7 {
		fmt.Println("Nice job")
	} else {
		fmt.Println("Not bad")
	}
}
