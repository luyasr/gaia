package log

var fuzzyStr = "***"

type FilterOption func(*Filter)

// FilterLevel is a FilterOption that sets the level of the Filter.
func FilterLevel(level Level) FilterOption {
	return func(f *Filter) {
		f.level = level
	}
}

// FilterKey is a FilterOption that sets the key of the Filter.
func FilterKey(key ...any) FilterOption {
	return func(f *Filter) {
		for _, v := range key {
			f.key[v] = struct{}{}
		}
	}
}

// FilterValue is a FilterOption that sets the value of the Filter.
func FilterValue(value ...any) FilterOption {
	return func(f *Filter) {
		for _, v := range value {
			f.value[v] = struct{}{}
		}
	}
}

func FilterFunc(filterFunc func(level Level, keyValues ...any) bool) FilterOption {
	return func(f *Filter) {
		f.filter = filterFunc
	}
}

type Filter struct {
	logger Logger
	level  Level
	key    map[any]struct{}
	value  map[any]struct{}
	filter func(level Level, keyValues ...any) bool
}

func NewFilter(logger Logger, opts ...FilterOption) *Filter {
	options := &Filter{
		logger: logger,
		key:    make(map[any]struct{}),
		value:  make(map[any]struct{}),
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

// Log print the kv pairs log.
func (f *Filter) Log(level Level, keyValues ...any) error {
	if level < f.level {
		return nil
	}

	if f.filter != nil && f.filter(level, keyValues...) {
		return nil
	}

	if len(f.key) > 0 || len(f.value) > 0 {
		for i := 0; i < len(keyValues); i += 2 {
			v := i + 1
			if v > len(keyValues) {
				continue
			}
			if _, ok := f.key[keyValues[i]]; ok {
				keyValues[v] = fuzzyStr
			}
			if _, ok := f.value[keyValues[i+1]]; ok {
				keyValues[v] = fuzzyStr
			}
		}
	}

	return f.logger.Log(level, keyValues...)
}
