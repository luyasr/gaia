package ioc

const (
	ConfigNamespace     = "config"
	DbNamespace         = "db"
	RouterNamespace     = "router"
	HandlerNamespace    = "handler"
	ControllerNamespace = "controller"
)

var (
	IocContainer = &Container{
		store: map[string]*ns{
			ConfigNamespace:     {name: ConfigNamespace, ioc: map[string]*ioc{}, priority: -9},
			DbNamespace:         {name: DbNamespace, ioc: map[string]*ioc{}, priority: -8},
			RouterNamespace:     {name: RouterNamespace, ioc: map[string]*ioc{}, priority: -7},
			HandlerNamespace:    {name: HandlerNamespace, ioc: map[string]*ioc{}, priority: -6},
			ControllerNamespace: {name: ControllerNamespace, ioc: map[string]*ioc{}, priority: -5},
		},
	}
)
