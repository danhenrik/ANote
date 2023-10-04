package httpAdapter

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Method  string
	Headers map[string][]string
	Body    string
	Query   map[string][]string
	Param   map[string][]string
	Path    string
	IP      string
	Cookies []*http.Cookie
}

func NewRequest(Method string,
	Headers map[string][]string,
	Body string,
	Query map[string][]string,
	Param map[string][]string,
	Path string,
	IP string,
	Cookies []*http.Cookie) Request {
	return Request{
		Method:  Method,
		Headers: Headers,
		Body:    Body,
		Query:   Query,
		Param:   Param,
		Path:    Path,
		IP:      IP,
		Cookies: Cookies,
	}
}

type Response struct {
	StatusCode int     `json:"status_code"` // HTTP status code
	Data       any     `json:"data"`        // Response data
	Message    *string `json:"message"`     // Error message
}

func NewErrorResponse(statusCode int, message string) Response {
	return Response{
		StatusCode: statusCode,
		Data:       nil,
		Message:    &message,
	}
}

func NewSuccessResponse(statusCode int, data any) Response {
	return Response{
		StatusCode: statusCode,
		Data:       data,
		Message:    nil,
	}
}

func NewNoContentRespone() Response {
	return Response{
		StatusCode: http.StatusNoContent,
		Data:       nil,
		Message:    nil,
	}
}

type Controller func(request Request) Response

func NewGinAdapter(c Controller) func(*gin.Context) {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		header := ctx.Request.Header
		requestBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Println("Error on read body", err)
			return
		}
		defer ctx.Request.Body.Close()
		body := string(requestBody)
		query := ctx.Request.URL.Query()
		param := make(map[string][]string)
		for _, entry := range ctx.Params {
			param[entry.Key] = append(param[entry.Key], entry.Value)
		}
		path := ctx.Request.URL.Path
		clientIp := ctx.ClientIP()
		cookies := ctx.Request.Cookies()

		request := NewRequest(
			method,
			header,
			body,
			query,
			param,
			path,
			clientIp,
			cookies,
		)

		response := c(request)
		ctx.JSON(response.StatusCode, response)
	}
}
