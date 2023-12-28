package ioc

var (
	apiHandlerContainer = &Container{store: map[string]Object{}}
)

func ApiHandler() *Container {
	return apiHandlerContainer
}
