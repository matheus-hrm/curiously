package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		return err
	}
	return nil
}

func WriteJson(c *gin.Context, code int, payload any) {
	c.JSON(code, payload)
}

func WriteError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}
