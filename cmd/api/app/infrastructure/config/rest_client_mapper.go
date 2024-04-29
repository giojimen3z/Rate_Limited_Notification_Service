//go:build !test
// +build !test

package config

import (
	"encoding/json"
	"reflect"

	"github.com/giojimen3z/Rate_Limited_Notification_Service/pkg/logger"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v8"
)

const (
	JSONParserError = "the body of the petition is incorrect, check and try again"
)

var (
	config   = &validator.Config{TagName: "validate"}
	validate = validator.New(config)
)

// MapRequestToStruct it receives customPool JSON object, expressed as an array of [] bytes
// and if it is correct it converts it into customPool struct otherwise it returns an error
func MapRequestToStruct(payload []byte, genericStruct interface{}) error {
	jsonErr := json.Unmarshal(payload, genericStruct)
	if jsonErr != nil {
		logger.Error("MapRequestToStruct - jsonErr "+jsonErr.Error(), jsonErr)
		return errors.New(JSONParserError)
	}

	value := reflect.ValueOf(genericStruct)
	indirectValue := reflect.Indirect(value)

	if indirectValue.Kind() == reflect.Slice || indirectValue.Kind() == reflect.Map {
		return nil
	}

	return validateStruct(genericStruct)
}

func validateStruct(genericStruct interface{}) error {
	return validate.Struct(genericStruct)
}
