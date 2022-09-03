package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gosrv/model"
	"net/http"
	"strings"
)

type Response struct {
	http.ResponseWriter
}

// Add CORS (Cross-Origin Resource Sharing) in response HTTP header.
// See more: https://developer.mozilla.org/en-US/docs/Glossary/CORS
func (res *Response) AddCors(origin string) {
	res.Header().Add(HEADER_ACCESS_CONTROL_ALLOW_ORIGIN, origin)
	res.Header().Add(HEADER_ACCESS_CONTROL_ALLOW_CREDENTIALS, "true")
	res.Header().Add(HEADER_ACCESS_CONTROL_ALLOW_HEADERS, strings.Join(
		[]string{
			HEADER_ACCEPT,
			HEADER_CONTENT_TYPE,
			HEADER_AUTHORIZATION,
		},
		", ",
	))
	res.Header().Add(HEADER_ACCESS_CONTROL_ALLOW_METHODS, strings.Join(
		[]string{
			METHOD_OPTIONS,
			METHOD_GET,
			METHOD_POST,
		},
		", ",
	))
}

// Response with error message in json data type
func (res *Response) responseWithErrorInfo(err error, statusCode int) {
	errBuffer, _ := json.Marshal(&map[string]string{
		"message": err.Error(),
	})

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Content-Length", fmt.Sprint(len(errBuffer)))
	res.WriteHeader(statusCode)
	res.Write(errBuffer)
}

// Response with json data type
func (res *Response) responseWithJsonData(data []byte, statusCode int) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Content-Length", fmt.Sprint(len(data)))
	res.WriteHeader(statusCode)
	res.Write(data)
}

func (res *Response) BadRequest(err error) {
	res.responseWithErrorInfo(err, http.StatusBadRequest)
}

func (res *Response) Unauthorized(err error) {
	res.responseWithErrorInfo(err, http.StatusUnauthorized)
}

func (res *Response) NotFound() {
	res.responseWithErrorInfo(fmt.Errorf("route not found"), http.StatusNotFound)
}

// Response internal server error status code to the client
func (res *Response) InternalServerError(err error) {
	res.responseWithErrorInfo(err, http.StatusInternalServerError)
}

// OK data (status code: 200).
// Use this template:
//
// {
//		"data": <buffer>
// }
func (res *Response) OkData(data interface{}) {
	dataBuffer, err := res.loadDataAsBuffer(res.buildPayload(data))
	if err != nil {
		res.InternalServerError(err)
		return
	}

	res.responseWithJsonData(dataBuffer, http.StatusOK)
}

// OK data with pagination data (status code: 200).
// Use this template:
//
// {
// 		"totalPages": <num>,
//		"data": [
//			<buffer>
//		]
// }
func (res *Response) OkDataWithPagination(data interface{}, totalPages uint) {
	payload := res.buildPayload(data)
	payload.TotalPages = totalPages

	dataBuffer, err := res.loadDataAsBuffer(payload)
	if err != nil {
		res.InternalServerError(err)
	}

	res.responseWithJsonData(dataBuffer, http.StatusOK)
}

// Created (status code: 201)
func (res *Response) Created(data interface{}) {
	dataBuffer, err := res.loadDataAsBuffer(data)
	if err != nil {
		res.InternalServerError(err)
	}

	res.responseWithJsonData(dataBuffer, http.StatusCreated)
}

// No Content (status code: 204)
func (res *Response) NoContent() {
	res.WriteHeader(http.StatusNoContent)
}

func (res *Response) loadDataAsBuffer(data interface{}) (buffer []byte, err error) {
	if !res.isDataBuffer(data) {
		buffer, err = json.Marshal(data)
	} else {
		buffer = data.([]byte)
	}
	return
}

func (res *Response) isDataBuffer(data interface{}) bool {
	return fmt.Sprintf("%T", data) == "[]uint8"
}

func (res *Response) buildPayload(data interface{}) *model.Data {
	version := md5.Sum([]byte(fmt.Sprintf("%+v", data)))

	payload := model.Data{}
	payload.Data = data
	payload.Version = hex.EncodeToString(version[:])

	return &payload
}
