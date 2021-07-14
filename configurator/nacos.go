package configurator

import (
	"sync"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type nacosInstance struct {
	conf   *NacosConfig
	client config_client.IConfigClient
}

// Key ...
func (n *nacosInstance) Key(group, key string) string {
	return ""
}

// Get ...
func (n *nacosInstance) Get(group, key string) (string, error) {
	return n.client.GetConfig(vo.ConfigParam{
		Group:  group,
		DataId: key,
	})
}

// Set ...
func (n *nacosInstance) Set(group, key, value string) (bool, error) {
	return n.client.PublishConfig(vo.ConfigParam{
		Group:   group,
		DataId:  key,
		Content: value,
	})
}

// Watch ...
func (n *nacosInstance) Watch(group, key string, cb func(key, value string)) error {
	return n.client.ListenConfig(vo.ConfigParam{
		Group:  group,
		DataId: key,
		OnChange: func(namespace, group, key, value string) {
			cb(key, value)
		},
	})
}

// --------------------------------

var (
	nacosOnce sync.Once
	nacos     *nacosInstance
)

// NewNacos ...
func NewNacos(confValue interface{}) *nacosInstance {
	nacosOnce.Do(func() {
		conf, ok := confValue.(*NacosConfig)
		if !ok {
			panic("invalid nacos config")
		}

		// TODO dir 配置

		nacosClient, err := clients.CreateConfigClient(map[string]interface{}{
			"serverConfigs": []constant.ServerConfig{
				{
					Scheme:      conf.Scheme,
					ContextPath: conf.ContextPath,
					IpAddr:      conf.IpAddr,
					Port:        conf.Port,
				},
			},
			"clientConfig": constant.ClientConfig{
				TimeoutMs:           10 * 1000,          // 请求超时时间, 单位毫秒
				NotLoadCacheAtStart: true,               // 启动时不加载缓存
				LogDir:              "/tmp/nacos/log",   // 日志目录
				CacheDir:            "/tmp/nacos/cache", // 缓存目录
				NamespaceId:         conf.Namespace,     // 命名空间
				Username:            conf.Username,      // 用户名
				Password:            conf.Password,      // 密码
			},
		})
		if err != nil {
			panic(err)
		}

		nacos = &nacosInstance{
			conf:   conf,
			client: nacosClient,
		}
	})
	return nacos
}
