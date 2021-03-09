package handlers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/xmarcoied/miauth/pkg"
)

const HardBodyLimit = 1024 * 64 // limit size of body

//BindTo bind a request body into given model
func BindTo(w http.ResponseWriter, r *http.Request, target interface{}) error {
	if err := render.DecodeJSON(http.MaxBytesReader(w, r.Body, HardBodyLimit), target); err != nil {
		RenderJSONError(w, r, http.StatusBadRequest, &pkg.Error{
			Code: pkg.ErrDecode,
			Msg:  "cannot decode a body",
		})
		return err
	}
	return nil
}

//Validation validates an entity
func Validation(w http.ResponseWriter, r *http.Request, entity interface{}) error {
	validate := validator.New()
	if err := validate.Struct(entity); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		t := reflect.TypeOf(entity)
		var listErrors []string

		for _, err := range validationErrors {
			field := err.Field()
			tag := err.Tag()
			f1, _ := t.FieldByName(field)
			jsonField, _ := f1.Tag.Lookup("json")
			listErrors = append(listErrors, fmt.Sprintf("Field %s is %s", jsonField, tag))
		}
		RenderJSONError(w, r, http.StatusBadRequest, &pkg.Error{
			Code: pkg.ErrRequestValidation,
			Msg:  strings.Join(listErrors, "\n"),
		})
		return err
	}
	return nil
}

//Encode64UserID encoding without padding a userID
func Encode64UserID(userID string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(userID))
}

//Decode64UserID encoding without padding a userID
func Decode64UserID(s string) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return "", errors.Wrap(err, "cannot decode userID")
	}
	return string(data), nil
}
