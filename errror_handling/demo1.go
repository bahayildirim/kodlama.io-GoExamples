package errrorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı")
		}

	} else {
		fmt.Println((f.Name()))
	}
}
