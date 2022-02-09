package main

import (
	"fmt"
)

type Product struct{
	id int
	name string
	stock int
	price int
}

type Warehouse struct{
	item []Product
}

type Store struct{
	name string
}

var allRecord []map[string][]*Product
// var RecordById map[int]*Product
var RecordByStore map[string][]*Product

func Record(s Store, p Product) []map[string][]*Product {
    // RecordById[p.id] = &p
    RecordByStore[s.name] = append(RecordByStore[s.name], &p)
	allRecord = append(allRecord, RecordByStore)
	// fmt.Println(s.name," nambah ",p.name,"dengan id",p.id)
	fmt.Println("dati function",allRecord)
	return allRecord
}

func (w *Warehouse)AddProduct(s Store, p Product) { 
	// fmt.Println("panjang item", len(w.item))
	if len(w.item) == 0 {
		p.id = 1
	}
	p.id = len(w.item) + 1
	w.item = append(w.item, p)
	// fmt.Println(w.item)
	Record(s,p)
	fmt.Println(s.name," nambah ",p.name,"dengan id",p.id)
} 

func (w *Warehouse)RemoveProduct(s Store, p Product) {
	result := make([]Product, len(w.item)-1)
	for _, v := range w.item {
		if v.name != p.name {
			result = append(result, v)
		}
	}
	w.item = result
	Record(s,p)
	fmt.Println(s.name," remove ",p.name,"dengan id",p.id)
}

func main() {
	fmt.Println("test dropezy")
	// RecordById = make(map[int]*Product)
    RecordByStore = make(map[string][]*Product)
	var (
	gudang Warehouse
	produk1 Product = Product{name: "Ciki",stock: 10,price: 10,}
	produk2 Product = Product{name: "Soda",stock: 20,price: 10,}
	produk3 Product = Product{name: "Minyak",stock: 30,price: 10,}
	store1 Store = Store{name: "jakarta 1",}
	store2 Store = Store{name: "jakarta 2",}
	store3 Store = Store{name: "jakarta 3",}
	)
	gudang.AddProduct(store1, produk1)
	fmt.Println("yang ada di gudang",gudang)
	gudang.AddProduct(store2, produk2)
	fmt.Println("yang ada di gudang",gudang)
	gudang.AddProduct(store3, produk3)
	fmt.Println("yang ada di gudang",gudang)
	gudang.RemoveProduct(store1, produk3)
	fmt.Println("yang ada di gudang",gudang)
	gudang.RemoveProduct(store2, produk1)
	fmt.Println("yang ada di gudang",gudang)
	fmt.Println("semua record ",allRecord)
	fmt.Println(len(gudang.item))
	fmt.Println(len(allRecord))
	// fmt.Println("semua record by id ",RecordById[2])
	for _, v := range RecordByStore["jakarta 1"] {
		fmt.Println("sampe",v)
	}
}