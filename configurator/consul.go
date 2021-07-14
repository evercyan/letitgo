package configurator

import (
	"fmt"
	"strings"
	"sync"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

type consulInstance struct {
	conf   *ConsulConfig
	client *api.Client
}

// Key /{group}/{key}
func (c *consulInstance) Key(group, key string) string {
	return fmt.Sprintf("%s/%s", strings.Trim(group, "/"), key)
}

// Get ...
func (c *consulInstance) Get(group, key string) (string, error) {
	v, _, err := c.client.KV().Get(c.Key(group, key), nil)
	if err != nil {
		return "", err
	}
	if v == nil {
		return "", nil
	}
	return string(v.Value), err
}

// Set ...
func (c *consulInstance) Set(group, key, value string) (bool, error) {
	_, err := c.client.KV().Put(&api.KVPair{
		Key:   c.Key(group, key),
		Flags: 0,
		Value: []byte(value),
	}, nil)
	return err == nil, err
}

// Watch ...
func (c *consulInstance) Watch(group, key string, cb func(key, value string)) error {
	key = c.Key(group, key)
	params := map[string]interface{}{
		"type": "key",
		"key":  key,
	}
	plan, err := watch.Parse(params)
	if err != nil {
		return err
	}

	plan.Handler = func(idx uint64, val interface{}) {
		if v, ok := val.(*api.KVPair); ok && v.Key == strings.TrimLeft(key, "/") {
			cb(key, string(v.Value))
		}
	}
	go plan.Run(fmt.Sprintf("%s:%d", c.conf.IpAddr, c.conf.Port))
	return nil
}

// --------------------------------

var (
	consulOnce sync.Once
	consul     *consulInstance
)

// NewConsul ...
func NewConsul(confValue interface{}) *consulInstance {
	consulOnce.Do(func() {
		conf, ok := confValue.(*ConsulConfig)
		if !ok {
			panic("invalid consul config")
		}
		config := api.DefaultConfig()
		config.Address = fmt.Sprintf("%s:%d", conf.IpAddr, conf.Port)
		consulClient, err := api.NewClient(config)
		if err != nil {
			panic(err)
		}

		consul = &consulInstance{
			conf:   conf,
			client: consulClient,
		}
	})
	return consul
}
