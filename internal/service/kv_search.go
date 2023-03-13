package service

import (
	"fmt"
	"kv-search-api/internal/model"
	"kv-search-api/internal/repository"
	"net/http"
	"sync"
)

type StorageService struct {
	StorageRepository *repository.Repository
	mu                sync.Mutex
}

func NewStorageService(s *repository.Repository) StorageService {
	return StorageService{StorageRepository: s}
}

func (s *StorageService) SetIfKeyNotUsedBefore(key string, value string) bool {
	if s.StorageRepository.IsKeyUsedBefore(key) || key == "" || value == "" {
		s.StorageRepository.SetPair(key, value)
		return true
	}
	return s.StorageRepository.SetPair(key, value)
}

func (s *StorageService) GetValue(key string) string {
	return s.StorageRepository.GetValue(key)
}

func (s *StorageService) GetRecords(req model.RecordsRequest) (model.RecordsResponse, error) {
	responseData, err := s.StorageRepository.GetRecords(req)
	var res model.RecordsResponse

	if err != nil {
		fmt.Println(err)
		res.Code = http.StatusInternalServerError
		res.Msg = "An error occured."
		res.Records = nil
		return res, err
	}

	res.Code = http.StatusNotFound
	res.Msg = "No data found."
	res.Records = responseData

	if len(responseData) > 0 {
		res.Code = 0
		res.Msg = "Success"
		res.Records = responseData
	}
	return res, nil
}
