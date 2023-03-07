package arrays

import "fmt"

func Demo3() {
	sayilar := [5]int{20, 30, 50, 10, 2}
	enbüyük := 0
	fmt.Println(sayilar)

	for i := 0; i < len(sayilar); i++ {
		if enbüyük < sayilar[i] {
			enbüyük = sayilar[i]
		}
	}
	fmt.Println(enbüyük)
}
