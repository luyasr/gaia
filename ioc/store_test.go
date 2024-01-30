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
	Container.Registry(DbNamespace, &mockIoc{name: "1"})
	Container.Registry(ConfigNamespace, &mockIoc{name: "2"})
	Container.Registry(ConfigNamespace, &mockIoc{name: "3"})
	Container.Registry(ControllerNamespace, &mockIoc{name: "4"})
	Container.Registry(RouterNamespace, &mockIoc{name: "5"})
	Container.RegistryNamespace("xxxxx")
	Container.Registry("xxxxx", &mockIoc{name: "8"}, 3)
	Container.Registry("xxxxx", &mockIoc{name: "10"}, 4)
	Container.Registry("xxxxx", &mockIoc{name: "12"}, 1)
	Container.Registry("xxxxx", &mockIoc{name: "14"}, 0)
	Container.Init()
	Container.Close()
}
