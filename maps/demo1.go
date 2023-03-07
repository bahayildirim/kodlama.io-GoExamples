package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"
	fmt.Println(sozluk["book"])
	deger, varMi := sozluk["lamp"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu: ", varMi)
}
