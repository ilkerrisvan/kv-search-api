package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"kv-search-api/internal/model"
	"kv-search-api/internal/service"
	"kv-search-api/pkg/util"
	"log"
	"net/http"
)

type StorageAPI struct {
	StorageService service.StorageService
}

func NewStorageAPI(s service.StorageService) StorageAPI {
	return StorageAPI{StorageService: s}
}

func (s StorageAPI) Create(w http.ResponseWriter, req *http.Request) {
	log.Printf("HTTP method: %s", req.Method)
	log.Printf("Endpoint: %s", req.URL)
	log.Printf("Request header: %s", req.Header)

	err := util.CheckHTTPMethod(req, http.MethodPost)
	if err != nil {
		Error(w, http.StatusBadRequest, err, err.Error())
		return
	}

	var respondMap model.KeyValueDbData
	body, _ := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &respondMap)
	if err != nil {
		e := errors.New("There is an error in the requested data. Check the data. Data should be JSON.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
		return
	}

	if s.StorageService.SetIfKeyNotUsedBefore(respondMap.Key, respondMap.Value) {
		e := errors.New("Key updated.")
		log.Printf("Error Message:%s", e.Error())
	}
	log.Printf("Respond Status Code: %d", 201)
	log.Printf("Respond Data: %s", respondMap)
	JSON(w, http.StatusCreated, respondMap)
}

func (s StorageAPI) Fetch(w http.ResponseWriter, req *http.Request) {
	log.Printf("HTTP method: %s", req.Method)
	log.Printf("Endpoint: %s", req.URL)
	log.Printf("Request header: %s", req.Header)

	err := util.CheckHTTPMethod(req, http.MethodGet)
	if err != nil {
		Error(w, http.StatusBadRequest, err, err.Error())
		return
	}

	key := req.URL.Query().Get("key")
	val := s.StorageService.GetValue(key)

	var respond model.KeyValueDbData
	if key == "" || val == "" {
		e := errors.New("Bad Request. The URL may be an incorrect or there may not be a value for the key value.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
	} else {
		respond.Key = key
		respond.Value = val
		log.Printf("Respond Status Code: %d", 200)
		log.Printf("Respond Data: %s", respond)
		JSON(w, http.StatusOK, respond)
	}
}

func (s StorageAPI) Search(w http.ResponseWriter, req *http.Request) {
	log.Printf("HTTP method: %s", req.Method)
	log.Printf("Endpoint: %s", req.URL)
	log.Printf("Request header: %s", req.Header)

	err := util.CheckHTTPMethod(req, http.MethodPost)
	if err != nil {
		Error(w, http.StatusBadRequest, err, err.Error())
		return
	}

	var recordsSearchRequest model.RecordsRequest
	var recordsSearchResponse model.RecordsResponse

	body, _ := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &recordsSearchRequest)
	if err != nil {
		log.Printf("An error occured with the json.")
	}
	recordsSearchResponse, err = s.StorageService.GetRecords(recordsSearchRequest)
	if err != nil {
		log.Printf("An error occured with the mongo query.")
	}

	if err != nil {
		e := errors.New("An error occured with the request.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		Error(w, http.StatusBadRequest, e, e.Error())
		return
	}

	log.Printf("Respond Status Code: %d", 201)
	JSON(w, http.StatusOK, recordsSearchResponse)
}
