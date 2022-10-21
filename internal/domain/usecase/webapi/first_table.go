package webapi

import "rest-api/internal/domain/entity"

type Service interface {
	GetAllService() []entity.FirstTable
}

type tableUsecase struct {
	service Service
}

func (u tableUsecase) GetAllUsecase() []entity.FirstTable {
	return u.service.GetAllService()
}
