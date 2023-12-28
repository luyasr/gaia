package ioc

import "github.com/gin-gonic/gin"

type Object interface {
	Init() error
	Name() string
}

type GinIRouter interface {
	Registry(r gin.IRouter)
}
