package middlewares

import (
	"fmt"
	"net/http"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var (
	validate     *validator.Validate
	validateOnce sync.Once
)

func initializeValidator() {
	validate = validator.New()
}

func ValidateRequest(validationStruct interface{}) gin.HandlerFunc {
	validateOnce.Do(func() {
		initializeValidator()
	})
	return func(c *gin.Context) {
		contentType := c.ContentType()

		if contentType == "application/json" {
			if err := validateJSON(c, validationStruct); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": formatValidationErrors(err, validationStruct)})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func validateJSON(c *gin.Context, validationStruct interface{}) error {
	if validationStruct == nil {
		return nil
	}

	if err := c.ShouldBindJSON(validationStruct); err != nil {
		return err
	}

	if err := validate.Struct(validationStruct); err != nil {
		return err
	}

	return nil
}

func formatValidationErrors(err error, validationStruct interface{}) map[string]string {
	errors := make(map[string]string)
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		t := reflect.TypeOf(validationStruct).Elem()
		for _, err := range validationErrs {
			field, _ := t.FieldByName(err.Field())
			fieldName := field.Tag.Get("json")
			errors[fieldName] = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fieldName, err.Tag())
		}
	}
	return errors
}
