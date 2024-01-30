package ioc

import (
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/log"
)

const defaultPriority = 0

type Container struct {
	store map[string]*ns
}

type ns struct {
	name     string
	ioc      map[string]*ioc
	priority int
}

type ioc struct {
	name     string
	object   Ioc
	priority int
}

func (c *Container) sort() []*ns {
	stored := make([]*ns, 0, len(c.store))
	for _, ns := range c.store {
		stored = append(stored, ns)
	}

	sort.Slice(stored, func(i, j int) bool {
		return stored[i].priority < stored[j].priority
	})

	return stored
}

func (n *ns) sort() []*ioc {
	sorted := make([]*ioc, 0, len(n.ioc))
	for _, ioc := range n.ioc {
		sorted = append(sorted, ioc)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].priority < sorted[j].priority
	})

	return sorted
}

func (c *Container) Init() error {
	stored := c.sort()
	for _, ns := range stored {
		sorted := ns.sort()
		for _, ioc := range sorted {
			err := ioc.object.Init()
			if err != nil {
				return err
			}
			log.Infof("%s init: %s", ns.name, ioc.name)
		}
	}

	return nil
}

func (c *Container) Get(namespace, name string) (Ioc, error) {
	if ns, exists := c.store[namespace]; exists {
		if ioc, exists := ns.ioc[name]; exists {
			return ioc.object, nil
		}
		return nil, errors.NotFound("ioc not found", "namespace: %s, name: %s", namespace, name)
	}

	return nil, errors.NotFound("namespace not found", "namespace: %s", namespace)
}

func (c *Container) RegistryNamespace(namespace string, priority ...int) {
	prio := defaultPriority
	if len(priority) > 0 {
		prio = priority[0]
	}

	if _, exists := c.store[namespace]; !exists {
		c.store[namespace] = &ns{name: namespace, ioc: map[string]*ioc{}, priority: prio}
		log.Infof("registry namespace: %s", namespace)
	}
}

func (c *Container) Registry(namespace string, object Ioc, priority ...int) {
	prio := defaultPriority
	if len(priority) > 0 {
		prio = priority[0]
	}
	objectName := object.Name()

	if ns, exists := c.store[namespace]; exists {
		ns.ioc[objectName] = &ioc{name: objectName, object: object, priority: prio}
		log.Infof("%s registry: %s", namespace, objectName)
	}
}

func (c *Container) GinIRouterRegistry(r gin.IRouter) {
	for _, ioc := range c.store[RouterNamespace].ioc {
		if router, ok := ioc.object.(GinIRouter); ok {
			router.Registry(r)
			log.Infof("GinIRouterRegistry: %s", ioc.name)
		}
	}
}
