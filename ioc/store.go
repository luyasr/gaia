package ioc

import (
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/log"
)

const defaultPriority = 0

type container struct {
	store  map[string]*ns
	sorted []*ns
	log    *log.Helper
}

type ns struct {
	name     string
	ioc      map[string]*ioc
	priority int
	sorted   []*ioc
}

type ioc struct {
	name     string
	object   Ioc
	priority int
}

func (c *container) SetLogger(l *log.Helper) {
	c.log = l
}

func (c *container) sort() {
	c.sorted = make([]*ns, 0, len(c.store))
	for _, ns := range c.store {
		c.sorted = append(c.sorted, ns)
	}

	sort.Slice(c.sorted, func(i, j int) bool {
		return c.sorted[i].priority < c.sorted[j].priority
	})
}

func (c *container) reverse() {
	length := len(c.sorted)
	for i := 0; i < length/2; i++ {
		c.sorted[i], c.sorted[length-i-1] = c.sorted[length-i-1], c.sorted[i]
	}
}

func (n *ns) sort() {
	n.sorted = make([]*ioc, 0, len(n.ioc))
	for _, ioc := range n.ioc {
		n.sorted = append(n.sorted, ioc)
	}

	sort.Slice(n.sorted, func(i, j int) bool {
		return n.sorted[i].priority < n.sorted[j].priority
	})
}

func (n *ns) reverse() {
	length := len(n.sorted)
	for i := 0; i < length/2; i++ {
		n.sorted[i], n.sorted[length-i-1] = n.sorted[length-i-1], n.sorted[i]
	}
}

func (c *container) Init() error {
	c.sort()
	for _, ns := range c.sorted {
		ns.sort()
		for _, ioc := range ns.sorted {
			err := ioc.object.Init()
			if err != nil {
				return err
			}
			c.log.Infof("%s init: %s", ns.name, ioc.name)
		}
	}

	return nil
}

func (c *container) Get(namespace, name string) (Ioc, error) {
	if ns, exists := c.store[namespace]; exists {
		if ioc, exists := ns.ioc[name]; exists {
			return ioc.object, nil
		}
		return nil, errors.NotFound("ioc not found", "namespace: %s, name: %s", namespace, name)
	}

	return nil, errors.NotFound("namespace not found", "namespace: %s, name: %s", namespace, name)
}

func (c *container) RegistryNamespace(namespace string, priority ...int) {
	if _, exists := c.store[namespace]; !exists {
		prio := defaultPriority
		if len(priority) > 0 {
			prio = priority[0]
		}

		c.store[namespace] = &ns{name: namespace, ioc: map[string]*ioc{}, priority: prio}
		c.log.Infof("registry namespace: %s", namespace)
	}
}

func (c *container) Registry(namespace string, object Ioc, priority ...int) {
	if ns, exists := c.store[namespace]; exists {
		prio := defaultPriority
		if len(priority) > 0 {
			prio = priority[0]
		}
		objectName := object.Name()

		ns.ioc[objectName] = &ioc{name: objectName, object: object, priority: prio}
		c.log.Infof("%s registry: %s", namespace, objectName)
	}
}

func (c *container) GinIRouterRegistry(r gin.IRouter) {
	for _, ioc := range c.store[RouterNamespace].ioc {
		if router, ok := ioc.object.(GinIRouter); ok {
			router.Registry(r)
			c.log.Infof("GinIRouterRegistry: %s", ioc.name)
		}
	}
}

// Close will close all the ioc.Closer
// 倒序关闭所有实现了 ioc.Closer 的对象
func (c *container) Close() error {
	c.reverse()
	for _, ns := range c.sorted {
		ns.reverse()
		for _, ioc := range ns.sorted {
			if closer, ok := ioc.object.(Closer); ok {
				if err := closer.Close(); err != nil {
					return err
				}
				c.log.Infof("%s close: %s", ns.name, ioc.name)
			}
		}
	}

	return nil
}
