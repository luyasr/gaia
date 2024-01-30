package ioc

import (
	"testing"
)

type mockIoc struct {
	name string
}

func (m *mockIoc) Init() error {
	return nil
}

func (m *mockIoc) Name() string {
	return m.name
}

func TestRegistryAndGet(t *testing.T) {
	IocContainer.Registry(DbNamespace, &mockIoc{})
	IocContainer.Registry(ConfigNamespace, &mockIoc{})
	IocContainer.Registry(ConfigNamespace, &mockIoc{})
	IocContainer.Registry(ControllerNamespace, &mockIoc{})
	IocContainer.Registry(RouterNamespace, &mockIoc{})
	IocContainer.RegistryNamespace("xxxxx")
	IocContainer.Registry("xxxxx", &mockIoc{name: "2"}, 3)
	IocContainer.Registry("xxxxx", &mockIoc{name: "4"}, 4)
	IocContainer.Registry("xxxxx", &mockIoc{name: "6"}, 1)
	IocContainer.Registry("xxxxx", &mockIoc{name: "8"}, 0)
	IocContainer.Init()
}
