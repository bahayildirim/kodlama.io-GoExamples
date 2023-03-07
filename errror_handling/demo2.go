package errrorhandling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklımdakiSayı := 50
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz.")
	}
	if tahmin > aklımdakiSayı {
		return "Daha küçük bir sayı giriniz", nil
	}
	if tahmin < aklımdakiSayı {
		return "Daha büyük bir sayı giriniz", nil
	}

	return "Bildiniz", nil
}

func Demo2() {
	mesaj, _ := TahminEt(49)
	fmt.Println(mesaj)
}
