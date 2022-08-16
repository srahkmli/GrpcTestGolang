package controller

import (
	"micro/api/pb/product"
	"micro/model"
)

func SampleRequestToProduct(r *product.SampleRequest) model.ProductModel {
	return model.ProductModel{
		Name: r.GetName(),
		Qty:  r.GetQty(),
	}
}

func PurchaseToSampleResponse(m model.PurchaseModel) *product.SampleResponse {
	return &product.SampleResponse{
		Data: m.Data,
	}
}
