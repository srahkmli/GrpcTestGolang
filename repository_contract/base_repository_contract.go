package repocontract

import "micro/model"

type IBaseRepository interface {
	StoreBaseModel(model.BaseModel1) error
	NotifySomeone(model.BaseModel2) error
}
