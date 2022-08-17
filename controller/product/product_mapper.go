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
func SampleRequestPoint(r *product.SamplePoint) model.PointModel {
	return model.PointModel{
		Point: r.GetPoint(),
	}
}

func PurchaseToSampleResponse(m model.PurchaseModel) *product.SampleResponse {
	return &product.SampleResponse{
		Data: m.Data,
	}
}
