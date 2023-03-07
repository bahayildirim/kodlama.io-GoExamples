package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1100

	if cekilmekIstenen > hesap {
		fmt.Print("Yetersiz Para")
	}
}
