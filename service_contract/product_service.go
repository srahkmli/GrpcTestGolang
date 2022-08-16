package servicecontract

import "micro/model"

type IProductService interface {
	Validate(string) bool
	Process(productModel model.ProductModel) (model.PurchaseModel, error)
}
