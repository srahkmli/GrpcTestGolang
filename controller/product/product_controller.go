package controller

import (
	"context"
	product "micro/api/pb/product"
	"micro/client/jtrace"
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

func (b *ProductController) SampleEndpoint(c context.Context, req *product.SampleRequest) (*product.SampleResponse, error) {
	span, _ := jtrace.T().SpanFromContext(c, "controller")
	defer span.Finish()
	zap.L().Info("an info level log")
	//if ok, violations := ValidateSampleRequest(req); !ok {
	//	return nil, gerrors.NewStatus(codes.Aborted).
	//		WithMessage("invalid name").
	//		AddBadRequest(violations...).
	//		AddFarsi("شما در وارد کردن اسم اشتباه کردید").
	//		MakeError()
	//}

	zap.L().Debug("a debug level log")

	reqModel := SampleRequestToProduct(req)

	resModel, err := b.productService.Process(c, reqModel)
	zap.L().Debug("last log")

	if err != nil {
		return nil, err
	}

	return PurchaseToSampleResponse(resModel), nil
}
