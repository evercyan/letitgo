package configurator

// 配置类型
const (
	TypeConsul = "consul"
	TypeNacos  = "nacos"
	TypeFile   = "file"
)

// NacosConfig ... nacos 配置
type NacosConfig struct {
	Scheme      string
	ContextPath string
	IpAddr      string
	Port        uint64
	Namespace   string
	Username    string
	Password    string
}

// ConsulConfig ... consul 配置
type ConsulConfig struct {
	IpAddr string
	Port   uint64
}

// FileConfig ... file 配置
type FileConfig struct {
	Path string
}
