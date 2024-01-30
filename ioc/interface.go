package ioc

import "github.com/gin-gonic/gin"

// Ioc is a interface for ioc
type Ioc interface {
	Init() error
	Name() string
}

// GinIRouter is a interface for gin.IRouter
type GinIRouter interface {
	Registry(r gin.IRouter)
}

// Closer is a interface for close
type Closer interface {
    Close() error
}
