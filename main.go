package main

import (
	"fmt"
	"golesson\variables"
)

func main() {
	variables.Demo1()
	//conditionals.Demo1()
	//var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//fmt.Println(functions.ToplaVariadic(list...))
	//ciftSayiCn := make(chan int)
	//tekSayiCn := make(chan int)
	//go channels.CiftSayilar(ciftSayiCn)
	//go channels.TekSayilar(tekSayiCn)
	//ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	//fmt.Println("Çarpım: ", ciftSayiToplam*tekSayiToplam)
	//product, _ := project.AddProduct()
	//fmt.Println(product)
	/*products, _ := project.GetAllProducts()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	} */
	for i := 1; i <= 10; i = i + 2 {
		fmt.Println("test")
	}
}
