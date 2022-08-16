package model

type ProductModel struct {
	Name string
	Qty  int32
}

func NewProductModel(name string, qty int32) ProductModel {
	return ProductModel{Name: name, Qty: qty}
}
