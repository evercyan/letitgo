package main

import (
	"fmt"
	"time"

	cfg "github.com/evercyan/letitgo/configurator"
)

var clients = map[string]func() cfg.Configurator{
	"nacos": func() cfg.Configurator {
		return cfg.NewConfigurator(cfg.TypeNacos, &cfg.NacosConfig{
			Scheme:      "http",
			ContextPath: "/nacos",
			IpAddr:      "127.0.0.1",
			Port:        8848,
			Namespace:   "local",
			Username:    "nacos",
			Password:    "nacos",
		})
	},
	"consul": func() cfg.Configurator {
		return cfg.NewConfigurator(cfg.TypeConsul, &cfg.ConsulConfig{
			IpAddr: "127.0.0.1",
			Port:   8500,
		})
	},
	"file": func() cfg.Configurator {
		return cfg.NewConfigurator(cfg.TypeFile, &cfg.FileConfig{
			Path: "./config.yaml",
		})
	},
}

func main() {
	name := "file" // consul, nacos, file

	group, key := "demo", "service"
	client := clients[name]()
	fmt.Println(client.Set(group, key, fmt.Sprintf("%s %s", name, time.Now().Format(time.Stamp))))
	fmt.Println(client.Get(group, key))
	client.Watch(group, key, func(key, value string) {
		fmt.Println("--------------------------------")
		fmt.Println(name, "watch", "key", key, "value", value)
		fmt.Println("--------------------------------")
	})

	for {
	}
}
