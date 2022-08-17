package service

import (
	"fmt"
	"go.uber.org/zap"
	"micro/model"
	repocontract "micro/repository_contract"
	"regexp"
)

type BaseService struct {
	baseRepository repocontract.IBaseRepository
}

func NewBaseService(repo repocontract.IBaseRepository) BaseService {
	return BaseService{
		baseRepository: repo,
	}
}

func (BaseService) Validate(userID string) bool {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	return isAlpha(userID)
}

func (b BaseService) Process(m model.BaseModel1) (model.BaseModel2, error) {
	result := model.BaseModel2{Data: fmt.Sprintf("Hello %s - %d", m.UserID, m.Code)}

	zap.L().Info("an process level log")
	if err := b.baseRepository.StoreBaseModel(m); err != nil {
		return result, err
	}
	if err := b.baseRepository.NotifySomeone(result); err != nil {
		return result, err
	}

	zap.L().Debug(fmt.Sprintf("result : %s ", result.Data))
	return result, nil
}
