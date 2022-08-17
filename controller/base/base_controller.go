package controller

import (
	"context"
	"micro/api/pb/base"
	"micro/client/jtrace"

	servicecontract "micro/service_contract"

	"go.uber.org/zap"
)

type BaseController struct {
	base.UnimplementedSampleAPIServer
	baseService servicecontract.IBaseService
}

func NewBaseController(service servicecontract.IBaseService) BaseController {
	return BaseController{
		baseService: service,
	}
}

func (b *BaseController) SampleEndpoint(c context.Context, req *base.SampleRequest) (*base.SampleResponse, error) {
	span, _ := jtrace.T().SpanFromContext(c, "controller")
	defer span.Finish()
	zap.L().Info("an info level log")
	//if ok, violations := ValidateSampleRequest(req); !ok {
	//	return nil, gerrors.NewStatus(codes.Aborted).
	//		WithMessage("invalid user ID").
	//		AddBadRequest(violations...).
	//		AddFarsi("شما در وارد کردن آیدی کاربر اشتباه کردید").
	//		MakeError()
	//}

	zap.L().Debug("a debug level log")

	model1 := SampleRequestToBaseModel1(req)
	model2, err := b.baseService.Process(model1)

	if err != nil {
		return nil, err
	}

	return BaseModel2ToSampleResponse(model2), nil
}
