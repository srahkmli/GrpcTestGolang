package model

type Product struct {
	Name string
	Qty  int32
}

func NewProduct(name string, qty int32) Product {
	return Product{Name: name, Qty: qty}
}
