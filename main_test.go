package main

import (
	"fmt"
	"testing"
)

var (
	store1 store = store{
		name: "jakarta barat",
	}
	store2 store = store{
		name: "jakarta timur",
	}
	product1 product = product{
		name: "shampo",
		stock: 1,
		price: 10,
	}
	product2 product = product{
		name: "sabun",
		stock: 1,
		price: 10,
	}
	gudang warehouse
)

func TestAddProduct(t *testing.T){
	expectation := product1
	var result product
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