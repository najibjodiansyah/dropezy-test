package main

import (
	"fmt"
	"testing"
)

var (
	store1 Store = Store{
		name: "jakarta barat",
	}
	store2 Store = Store{
		name: "jakarta timur",
	}
	product1 Product = Product{
		name: "shampo",
		stock: 1,
		price: 10,
	}
	product2 Product = Product{
		name: "sabun",
		stock: 1,
		price: 10,
	}
	gudang Warehouse
)

func TestAddProduct(t *testing.T){
	expectation := product1
	var result Product
	gudang.AddProduct(store1, product1)
	gudang.AddProduct(store2, product2)
	for _, v := range gudang.item {
		if v.name == product1.name{
			result = product1
		}
	}
	equal := result
	fmt.Println(expectation,equal,gudang)
	if equal != expectation {
		t.Errorf("Expected %v but got %v", expectation, equal)
	}
}

func TestRemoveProduct(t *testing.T){
	gudang.RemoveProduct(store1, product1)
	expectation := true
	equal := false
	for _, v := range gudang.item {
		if v.name != product1.name{
			equal = true
		}
	}
	if equal != expectation {
		t.Errorf("Expected %v but got %v", expectation, equal)
	}
}