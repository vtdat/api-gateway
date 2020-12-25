package controllers

import "github.com/gin-gonic/gin"

type AppResponse struct {
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
	Desc    string      `json:"description"`
}

func Response(code int, message interface{}, desc string) AppResponse {
	return AppResponse{Code: code, Message: message, Desc: desc}
}

func Abort(c *gin.Context, statusCode int, msg interface{}, desc string) {
	c.AbortWithStatusJSON(statusCode, Response(statusCode, msg, desc))
}
