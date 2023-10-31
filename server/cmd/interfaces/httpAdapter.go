package httpAdapter

import (
	"anote/internal/errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// =======================================================================================================================
// =================================== Types =============================================================================
// =======================================================================================================================

type UserIdentity struct {
	ID    string
	Email string
}

type Request struct {
	Method      string
	Headers     map[string][]string
	Body        string
	QueryParams map[string][]string
	PathParams  map[string][]string
	Path        string
	IP          string
	Cookies     []*http.Cookie
	User        UserIdentity
	FileHeaders map[string][]*multipart.FileHeader
	Files       []string
	Raw         *gin.Context
}

func NewRequest(
	Method string,
	Headers map[string][]string,
	Body string,
	QueryParams map[string][]string,
	PathParams map[string][]string,
	Path string,
	IP string,
	Cookies []*http.Cookie,
	FileHeaders map[string][]*multipart.FileHeader,
	Files []string,
	ctx *gin.Context,
) Request {
	return Request{
		Method:      Method,
		Headers:     Headers,
		Body:        Body,
		QueryParams: QueryParams,
		PathParams:  PathParams,
		Path:        Path,
		IP:          IP,
		Cookies:     Cookies,
		FileHeaders: FileHeaders,
		Files:       Files,
		Raw:         ctx,
	}
}

func (req *Request) GetHeader(key string) (string, bool) {
	if value, ok := req.Headers[key]; ok && len(value) > 0 {
		return value[0], true
	}
	return "", false
}

func (req *Request) GetSingleQuery(key string) (string, bool) {
	if value, ok := req.QueryParams[key]; ok && len(value) > 0 {
		return value[0], true
	}
	return "", false
}

func (req *Request) GetQuerySlice(key string) ([]string, bool) {
	if value, ok := req.QueryParams[key]; ok && len(value) > 0 {
		return value, true
	}
	return nil, false
}

func (req *Request) GetSingleParam(key string) (string, bool) {
	if value, ok := req.PathParams[key]; ok && len(value) > 0 {
		return value[0], true
	}
	return "", false
}

// =======================================================================================================================

type Response struct {
	StatusCode uint    `json:"-"`       // HTTP status code
	Data       any     `json:"data"`    // Response data
	Message    *string `json:"message"` // Error message
}

func NewErrorResponse(statusCode uint, message string) Response {
	return Response{
		StatusCode: statusCode,
		Data:       nil,
		Message:    &message,
	}
}

func NewSuccessResponse(statusCode uint, data any) Response {
	return Response{
		StatusCode: statusCode,
		Data:       data,
		Message:    nil,
	}
}

func NewNoContentResponse() Response {
	return Response{
		StatusCode: http.StatusNoContent,
		Data:       nil,
		Message:    nil,
	}
}

// =======================================================================================================================
// =================================== HTTP Adapter ======================================================================
// =======================================================================================================================

type Controller func(request Request) Response
type Middleware func(Request) (Request, *errors.AppError)

// convert gin request into our app request
func NewGinAdapter(
	c Controller,
	middlewares ...Middleware,
) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// create request object
		method := ctx.Request.Method
		header := ctx.Request.Header
		requestBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Println("Error on read body", err)
			return
		}
		defer ctx.Request.Body.Close()
		body := string(requestBody)
		queryParams := ctx.Request.URL.Query()
		pathParams := make(map[string][]string)
		for _, entry := range ctx.Params {
			pathParams[entry.Key] = append(pathParams[entry.Key], entry.Value)
		}
		path := ctx.Request.URL.Path
		clientIp := ctx.ClientIP()
		cookies := ctx.Request.Cookies()

		request := NewRequest(
			method,
			header,
			body,
			queryParams,
			pathParams,
			path,
			clientIp,
			cookies,
			nil,
			nil,
			ctx,
		)

		// execute middlewares
		for _, middleware := range middlewares {
			newReq, err := middleware(request)
			if err != nil {
				ctx.JSON(int(err.Status), err.Message)
				return
			}
			request = newReq
		}

		// execute controller/handler
		response := c(request)

		// parse response and return
		ctx.JSON(int(response.StatusCode), response)

		if len(request.Files) > 0 {
			for _, file := range request.Files {
				// delete from temp folder
				path := "internal/tmp/" + file
				if _, err := os.Stat(path); err == nil {
					log.Println("[httpAdapter] Tmp file", file, "exists, deleting it...")
					err := os.Remove(path)
					if err != nil {
						log.Println("[httpAdapter] Error on delete file", err)
					}
					log.Println("[httpAdapter] Tmp file", file, "successfully deleted")
				}
			}
		}
	}
}
