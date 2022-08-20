package controller

import (
	"context"
	"fmt"
	product "micro/api/pb/product"
	"micro/client/jtrace"
	"micro/model"
	servicecontract "micro/service_contract"

	"go.uber.org/zap"
)

type ProductController struct {
	product.UnimplementedSampleAPIServer
	productService servicecontract.IProductService
}

func NewProductController(service servicecontract.IProductService) ProductController {
	return ProductController{
		productService: service,
	}
}

func (b *ProductController) SampleEndpointSet(c context.Context, req *product.SampleRequest) (*product.SampleResponse, error) {
	span, c := jtrace.T().SpanFromContext(c, "controller[SET]")
	defer span.Finish()
	zap.L().Info("an info level log")

	reqModel := SampleRequestToProduct(req)
	err := b.productService.SetProcess(c, reqModel)
	if err != nil {
		return nil, err
	}

	zap.L().Debug("last log")
	sr := model.PurchaseModel{Data: fmt.Sprintf("Hello %s - %d", reqModel.Name, reqModel.Qty)}
	return PurchaseToSampleResponse(sr), nil
}
func (b *ProductController) SampleEndpointGet(c context.Context, req *product.SamplePoint) (*product.SampleResponse, error) {
	span, c := jtrace.T().SpanFromContext(c, "controller[GET]")
	defer span.Finish()
	zap.L().Info("an info level log")

	reqModel := SampleRequestPoint(req)
	resModel, err := b.productService.GetProcess(c, reqModel)
	if err != nil {
		return nil, err
	}
	zap.L().Debug("last log")
	sr := model.PurchaseModel{Data: fmt.Sprintf("Hello %s - %d", resModel.Name, resModel.Qty)}

	return PurchaseToSampleResponse(sr), nil
}
