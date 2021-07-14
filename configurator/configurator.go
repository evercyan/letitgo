package configurator

// Configurator ... 配置器接口
type Configurator interface {
	Key(group, key string) string
	Get(group, key string) (string, error)
	Set(group, key, value string) (bool, error)
	Watch(group, key string, cb func(key, value string)) error
	// TODO 服务注册发现
}

// NewConfigurator ...
func NewConfigurator(confType string, confValue interface{}) Configurator {
	switch confType {
	case TypeNacos:
		return NewNacos(confValue)
	case TypeConsul:
		return NewConsul(confValue)
	case TypeFile:
		return NewFile(confValue)
	default:
		panic("invalid configurator confType")
	}
}
