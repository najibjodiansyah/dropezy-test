package main

import "testing"

var (
	product1 product = product{
		name: "shampo",
		stock: 1,
		price: 10,
	}
	gudang warehouse
)

func TestAddProduct(t *testing.T){
	expectation := product1
	gudang.AddProduct(product1)
	equal := gudang.item[0]
	if equal != expectation {
		t.Errorf("Expected %v but got %v", expectation, equal)
	}
}