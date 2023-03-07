package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Engin"
	fmt.Println(s.Count(isim, "n"))
	fmt.Println(s.Contains(isim, "n"))
}
