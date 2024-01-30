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

func (m *mockIoc) Close() error {
	return nil
}

func TestRegistryAndGet(t *testing.T) {
	IocContainer.Registry(DbNamespace, &mockIoc{name: "1"})
	IocContainer.Registry(ConfigNamespace, &mockIoc{name: "2"})
	IocContainer.Registry(ConfigNamespace, &mockIoc{name: "3"})
	IocContainer.Registry(ControllerNamespace, &mockIoc{name: "4"})
	IocContainer.Registry(RouterNamespace, &mockIoc{name: "5"})
	IocContainer.RegistryNamespace("xxxxx")
	IocContainer.Registry("xxxxx", &mockIoc{name: "8"}, 3)
	IocContainer.Registry("xxxxx", &mockIoc{name: "10"}, 4)
	IocContainer.Registry("xxxxx", &mockIoc{name: "12"}, 1)
	IocContainer.Registry("xxxxx", &mockIoc{name: "14"}, 0)
	IocContainer.Init()
	IocContainer.Close()
}
