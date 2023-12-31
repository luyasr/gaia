package ioc

import "github.com/gin-gonic/gin"

type Container struct {
	store map[string]Object
}

func (c *Container) Init() error {
	for _, obj := range c.store {
		if err := obj.Init(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Container) Get(name string) Object {
	return c.store[name]
}

func (c *Container) Registry(object Object) {
	if _, exists := c.store[object.Name()]; !exists {
		c.store[object.Name()] = object
	}
}

func (c *Container) GinIRouterRegistry(r gin.IRouter) {
	for _, obj := range c.store {
		if router, ok := obj.(GinIRouter); ok {
			router.Registry(r)
		}
	}
}
