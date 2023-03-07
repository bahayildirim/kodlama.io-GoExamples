package deferstatement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "Çift sayıdır"
	}
	if sayi%2 == 1 {
		return "Tek sayıdır"
	}
	return "Belli değil"
}

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
