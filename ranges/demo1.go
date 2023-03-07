package ranges

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	for i, sehir := range sehirler {
		fmt.Print(i, " ")
		fmt.Println(sehir)
	}
}
