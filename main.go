package main

import "fmt"

type product struct{
	name string
	stock int
	price int
}

type warehouse struct{
	item []product
}

type store struct{
	name string
}

func (w *warehouse)AddProduct(p product) {
	 w.item = append(w.item, p)
	 fmt.Println("add")
} 

func (w *warehouse)RemoveProduct(p product) {
	result := make([]product, len(w.item)-1)
	for _, v := range w.item {
		if v.name != p.name {
			result = append(result, v)
		}
	}
	w.item = result
}

func main() {
	
	fmt.Println("test dropezy")
	var gudang warehouse
	var produk1 product = product{
		name: "Ciki",
		stock: 1,
		price: 1,
	}
	var produk2 product = product{
		name: "Soda",
		stock: 2,
		price: 1,
	}
	var produk3 product = product{
		name: "Minyak",
		stock: 3,
		price: 1,
	}
	gudang.AddProduct(produk1)
	fmt.Println("yang ada di gudang",gudang.item)
	gudang.AddProduct(produk2)
	gudang.AddProduct(produk3)
	fmt.Println("yang ada di gudang",gudang.item)

}