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
			ConfigNamespace:     {name: ConfigNamespace, ioc: map[string]*ioc{}, priority: -10},
			DbNamespace:         {name: DbNamespace, ioc: map[string]*ioc{}, priority: -9},
			RouterNamespace:     {name: RouterNamespace, ioc: map[string]*ioc{}, priority: -8},
			HandlerNamespace:    {name: HandlerNamespace, ioc: map[string]*ioc{}, priority: -7},
			ControllerNamespace: {name: ControllerNamespace, ioc: map[string]*ioc{}, priority: -6},
		},
	}
)
