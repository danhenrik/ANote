package middlewares

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/errors"
	"encoding/json"
	"log"
	"mime/multipart"
	"strings"
)

func ParseMultipart(req httpAdapter.Request) (httpAdapter.Request, *errors.AppError) {
	log.Println("[ParseMultipart Middleware] Parsing multipart")
	header, _ := req.Headers["Content-Type"]
	boundary := strings.Split(strings.Split(header[0], ";")[1], "=")[1]
	reader := multipart.NewReader(strings.NewReader(req.Body), boundary)

	form, err := reader.ReadForm(10 << 20)
	if err != nil {
		log.Println("[ParseMultipart Middleware] Error on read form:", err)
		return req, errors.NewAppError(400, "Invalid content-type")
	}

	// convert form to plain map[string]string
	planefiedForm := make(map[string]string)
	for key, value := range form.Value {
		planefiedForm[key] = value[0]
	}
	jsonString, err := json.Marshal(planefiedForm)
	log.Println("[ParseMultipart Middleware] JSON Created:", string(jsonString))
	if err != nil {
		log.Println("[ParseMultipart Middleware] Error on read form:", err)
		return req, errors.NewAppError(400, "Invalid content-type")
	}
	req.FileHeaders = form.File
	req.Body = string(jsonString)
	return req, nil
}
