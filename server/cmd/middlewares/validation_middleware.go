package middlewares

import (
	"net/http"
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
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// Function to validate the JSON body
func validateJSON(c *gin.Context, validationStruct interface{}) error {
	if err := c.ShouldBindJSON(validationStruct); err != nil {
		return err
	}

	if err := validate.Struct(validationStruct); err != nil {
		return err
	}

	return nil
}
