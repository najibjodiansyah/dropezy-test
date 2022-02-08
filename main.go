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

func (w *warehouse)AddProduct(s store, p product) {
	 w.item = append(w.item, p)
} 

func (w *warehouse)RemoveProduct(s store, p product) {
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
	var (
	produk1 product = product{
		name: "Ciki",
		stock: 1,
		price: 1,
	}
	produk2 product = product{
		name: "Soda",
		stock: 2,
		price: 1,
	}
	produk3 product = product{
		name: "Minyak",
		stock: 3,
		price: 1,
	}
	store1 store = store{
		name: "jakarta barat",
	}
	store2 store = store{
		name: "jakarta timur",
	}
	store3 store = store{
		name: "jakarta selatan",
	}
	)
	gudang.AddProduct(store1, produk1)
	fmt.Println("yang ada di gudang",gudang.item)
	gudang.AddProduct(store2, produk2)
	gudang.AddProduct(store3, produk3)
	fmt.Println("yang ada di gudang",gudang.item)

}