package util

import (
	"errors"
	"log"
	"net/http"
)

func CheckHTTPMethod(req *http.Request, method string) error {
	if req.Method != method {
		e := errors.New("Request method is not allowed.")
		log.Printf("Error Message:%s", e.Error())
		log.Printf("Respond Status Code: %d", 400)
		return e
	}
	return nil
}


