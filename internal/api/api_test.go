package api

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"kv-search-api/internal/model"
	"kv-search-api/internal/repository"
	"kv-search-api/internal/service"
	"kv-search-api/pkg/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func apiSetup(mongoClient *mongo.Client) StorageAPI {
	storageRepository := repository.NewRepository(mongoClient)
	storageService := service.NewStorageService(storageRepository)
	storageAPI := NewStorageAPI(storageService)
	return storageAPI
}

func TestCreate(t *testing.T) {
	w := httptest.NewRecorder()
	mongoClient := config.DBConnect()
	api := apiSetup(mongoClient)
	respondMapTest := model.KeyValueDbData{
		Key:   "testKeyData",
		Value: "testValueData",
	}
	body, _ := json.Marshal(respondMapTest)
	r := httptest.NewRequest("POST", "/api/create", bytes.NewReader(body))
	api.Create(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)

	var testRespondMap model.KeyValueDbData
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&testRespondMap); err != nil {
		t.Error(err)
	}
	assert.Equal(t, respondMapTest, testRespondMap)
}

func TestFetch(t *testing.T) {
	w := httptest.NewRecorder()
	mongoClient := config.DBConnect()
	api := apiSetup(mongoClient)

	respondMapTest := model.KeyValueDbData{
		Key:   "testKeyData",
		Value: "testValueData",
	}
	body, _ := json.Marshal(respondMapTest)
	r := httptest.NewRequest("POST", "/api/create", bytes.NewReader(body))
	api.Create(w, r)

	responseFetch := httptest.NewRequest("GET", "/api/fetch?key=testKeyData", nil)
	api.Fetch(w, responseFetch)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestSearch(t *testing.T) {
	w := httptest.NewRecorder()
	mongoClient := config.DBConnect()
	api := apiSetup(mongoClient)

	recordsSearchRequestTest := model.RecordsRequest{
		StartDate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  2700,
		MaxCount:  3000,
	}
	body, _ := json.Marshal(recordsSearchRequestTest)
	r := httptest.NewRequest("POST", "/api/create", bytes.NewReader(body))
	api.Search(w, r)

	recordsSearchResponseTest := model.RecordsResponse{
		Code: 0,
		Msg:  "Success",
	}

	var recordsSearchResponse model.RecordsResponse
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&recordsSearchResponse); err != nil {
		t.Error(err)
	}
	assert.Equal(t, recordsSearchResponseTest.Msg, recordsSearchResponse.Msg)
}
