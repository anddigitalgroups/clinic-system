package api

import "github.com/gin-gonic/gin"

type APIError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func RespondWithError(c *gin.Context, code int, errCode string, message string) {
	c.AbortWithStatusJSON(code, APIError{
		Error:   errCode,
		Message: message,
	})
}
