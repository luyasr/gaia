package ioc

import (
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
)

const (
	ConfigNamespace     = "config"
	DbNamespace         = "db"
	HandlerNamespace    = "handler"
	ControllerNamespace = "controller"
	userPriority        = 0
)

var (
	Container = &container{
		store: map[string]*ns{
			ConfigNamespace:     {name: ConfigNamespace, ioc: map[string]*ioc{}, priority: -9},
			DbNamespace:         {name: DbNamespace, ioc: map[string]*ioc{}, priority: -8},
			HandlerNamespace:    {name: HandlerNamespace, ioc: map[string]*ioc{}, priority: -7},
			ControllerNamespace: {name: ControllerNamespace, ioc: map[string]*ioc{}, priority: -6},
		},
	}
)

func init() {
	Container.SetLogger(log.NewHelper(zerolog.DefaultLogger))
}
