package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/laidingqing/stackbuild/handler/api/errors"
)

// ErrorCode writes the json-encoded error message to the response.
func ErrorCode(w http.ResponseWriter, err error, status int) {
	JSON(w, &errors.Error{Message: err.Error()}, status)
}

// InternalError writes the json-encoded error message
func InternalError(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 500)
}

// InternalErrorf writes the json-encoded error message
func InternalErrorf(w http.ResponseWriter, format string, a ...interface{}) {
	ErrorCode(w, fmt.Errorf(format, a...), 500)
}

// NotImplemented writes the json-encoded error message
func NotImplemented(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 501)
}

// NotFound writes the json-encoded error message
func NotFound(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 404)
}

// NotFoundf writes the json-encoded error message
func NotFoundf(w http.ResponseWriter, format string, a ...interface{}) {
	ErrorCode(w, fmt.Errorf(format, a...), 404)
}

// Unauthorized writes the json-encoded error message
func Unauthorized(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 401)
}

// Forbidden writes the json-encoded error message
func Forbidden(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 403)
}

// BadRequest writes the json-encoded error message
func BadRequest(w http.ResponseWriter, err error) {
	ErrorCode(w, err, 400)
}

// BadRequestf writes the json-encoded error message
func BadRequestf(w http.ResponseWriter, format string, a ...interface{}) {
	ErrorCode(w, fmt.Errorf(format, a...), 400)
}

// JSON writes
func JSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.Encode(v)
}
