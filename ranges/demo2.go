package ranges

import "fmt"

func Demo2() {
	total := 0
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		if num%2 == 1 {
			fmt.Println(num)
			total += num
		}
	}
	fmt.Println(total)
}
