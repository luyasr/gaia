package ioc

// apiHandlerContainer is a singleton container for API handlers.
var (
	apiHandlerContainer = &Container{store: map[string]Ioc{}}
)

// ApiHandler returns the singleton container for API handlers.
func ApiHandler() *Container {
	return apiHandlerContainer
}