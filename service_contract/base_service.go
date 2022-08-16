package servicecontract

import "micro/model"

type IBaseService interface {
	Validate(string) bool
	Process(model.BaseModel1) (model.BaseModel2, error)
}
