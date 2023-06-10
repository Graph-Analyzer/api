package _interface

import "github.com/gin-gonic/gin"

type Action interface {
	Invoke(context *gin.Context)
}
