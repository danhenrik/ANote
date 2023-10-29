package middlewares

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/errors"
	"anote/internal/helpers"
	"log"
	"os"
	"slices"
	"strings"
)

func SaveFile(fileField string, allowedExtensions []string) httpAdapter.Middleware {
	return func(req httpAdapter.Request) (httpAdapter.Request, *errors.AppError) {
		for _, fHeader := range req.FileHeaders[fileField] {
			file, err := fHeader.Open()
			if err != nil {
				log.Println("[SaveFile Middleware] Error on read file:", err)
				return req, errors.NewAppError(400, "Invalid content-type")
			}
			defer file.Close()

			splitted := strings.Split(fHeader.Filename, ".")
			ext := splitted[len(splitted)-1]
			if !slices.Contains(allowedExtensions, ext) {
				log.Println("[SaveFile Middleware] Error on save file: invalid extension")
				return req, errors.NewAppError(400, "Invalid file extension")
			}

			uuid := helpers.NewUUID()
			filename := uuid + "." + ext

			log.Println("[SaveFile Middleware] Saving file", filename)
			f, err := os.OpenFile("internal/tmp/"+filename, os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				log.Println("[SaveFile Middleware] Error on save file, proceding anyways. Err:", err)
				return req, nil
			}
			defer f.Close()

			fileContent := make([]byte, fHeader.Size)
			file.Read(fileContent)
			f.Write(fileContent)

			req.Files = append(req.Files, filename)
		}
		return req, nil
	}
}
