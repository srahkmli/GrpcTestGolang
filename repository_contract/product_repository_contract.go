package repocontract

import "micro/model"

type IProductRepository interface {
	StoreProductModel(model.ProductModel) error
	NotifyPurchase(model.PurchaseModel) error
}
