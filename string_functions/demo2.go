package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Engin"
	fmt.Println(s.HasPrefix(isim, "Eng"))
	fmt.Println(s.HasPrefix(isim, "in"))
	harfler := []string{"E", "n", "g", "i", "n"}
	fmt.Println(s.Join(harfler, " "))
	fmt.Println(s.Replace(s.Join(harfler, " "), " ", "-", -1))
}
