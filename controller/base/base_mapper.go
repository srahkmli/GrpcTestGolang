package controller

import (
	"micro/api/pb/base"
	"micro/model"
)

func SampleRequestToBaseModel1(r *base.SampleRequest) model.BaseModel1 {
	return model.BaseModel1{
		UserID: r.GetUserID(),
		Code:   r.GetData(),
	}
}

func BaseModel2ToSampleResponse(m model.BaseModel2) *base.SampleResponse {
	return &base.SampleResponse{
		Data: m.Data,
	}
}
