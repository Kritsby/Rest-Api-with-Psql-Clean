package service

import "rest-api/internal/domain/entity"

type FirstTableStorage interface {
	GetAllAdapter() []entity.FirstTable
}

type firstTableService struct {
	storage FirstTableStorage
}

func NewTableService(storage FirstTableStorage) *firstTableService {
	return &firstTableService{storage: storage}
}

func (s firstTableService) GetAllService() []entity.FirstTable {
	return s.storage.GetAllAdapter()
}
