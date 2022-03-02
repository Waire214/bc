package helper

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	INVALID_MESSAGE_ERROR   = "The message format read from the given topic is invalid"
	VALIDATION_ERROR        = "The request has validation errors"
	REQUEST_NOT_FOUND       = "The requested resource was NOT found"
	DUPLICATE_REQUEST_ERROR = "A resource having the same identifier already exist	"
	GENERIC_ERROR           = "Generic error occurred. See stacktrace for details"
	AUTHORIZATION_ERROR     = "You do NOT have adequate permission to access this resource"
	NO_PRINCIPAL            = "Principal identifier NOT provided"
	MongoDBError            = "MONGO_DB_ERROR"
	NoResourceFound         = "this resource does not exist"
	NoRecordFound           = "sorry. no record found"
	NoErrorsFound           = "no errors at the moment"
)

func (err ErrorResponse) Error() string {
	var errorBody ErrorBody
	return fmt.Sprintf("%v", errorBody)
}

func ErrorArrayToError(errorBody []validator.FieldError) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()

	for _, value := range errorBody {
		body := ErrorBody{Code: VALIDATION_ERROR, Source: Config.AppName, Message: value.Error()}
		errorResponse.Errors = append(errorResponse.Errors, body)
	}
	return errorResponse
}

func ErrorMessage(code string, message string) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()
	errorResponse.Errors = append(errorResponse.Errors, ErrorBody{Code: code, Source: "buy-coin-users", Message: message})

	return errorResponse
}
