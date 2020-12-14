package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details`
}

// codes 是用來避免建立重複代碼的 Error
var codes = map[int]string{}

// NewError 可以用來建立 Error struct
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("錯誤 %d 已經存在，請更換一個", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

// Error 可以顯示錯誤訊息
func (e *Error) Error() string {
	return fmt.Sprintf("錯誤 %d, 錯誤訊息：%s", e.Code(), e.Msg())
}

// Code 會顯示系統中的錯誤代碼
func (e *Error) Code() int {
	return e.code
}

// Msg 會顯示錯誤訊息
func (e *Error) Msg() string {
	return e.msg
}

// Msgf 會顯示格式化後的錯誤訊息
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

// Details 會顯示錯誤的詳細資料
func (e *Error) Details() []string {
	return e.details
}

// WithDetails 可以添加錯誤訊息的詳細資料
func (e *Error) WithDetails(details ...string) *Error {
	// NOTICE: 為什麼這裡要用把 pointer 解開？
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

// StatusCode 會將 Error 轉換成 httpStatusCode
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequest.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
