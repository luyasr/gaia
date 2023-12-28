package ioc

import "github.com/gin-gonic/gin"

type Container struct {
	store map[string]Object
}

// 加载所有对象
func (c *Container) Load() error {
	for _, obj := range c.store {
		obj := obj
		if err := obj.Init(); err != nil {
			return err
		}
	}

	return nil
}

// 获取对象
func (c *Container) Get(name string) Object {
	return c.store[name]
}

func (c *Container) Registry(object Object) {
	if obj, exists := c.store[object.Name()]; !exists {
		c.store[object.Name()] = obj
	}
}

// 注册所有路由
func (c *Container) GinIRouterRegistry(r gin.IRouter) {
	for _, obj := range c.store {
		if router, ok := obj.(GinIRouter); ok {
			router.Registry(r)
		}
	}
}
