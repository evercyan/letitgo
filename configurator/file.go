package configurator

import (
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type fileInstance struct {
	conf      *FileConfig
	client    *viper.Viper
	watchOnce sync.Once
	watchList []watchInfo
}

type watchInfo struct {
	key   string                  // key
	value string                  // 旧值
	cb    func(key, value string) // 回调函数
}

// Key /{group}/{key}
func (f *fileInstance) Key(group, key string) string {
	return strings.Trim(fmt.Sprintf("%s.%s", group, key), ".")
}

// Get ...
func (f *fileInstance) Get(group, key string) (string, error) {
	return fmt.Sprintf("%v", f.client.Get(f.Key(group, key))), nil
}

// Set ...
func (f *fileInstance) Set(group, key, value string) (bool, error) {
	return true, nil
}

// Watch ...
func (f *fileInstance) Watch(group, key string, cb func(key, value string)) error {
	f.watchOnce.Do(func() {
		f.client.WatchConfig()
		f.client.OnConfigChange(func(e fsnotify.Event) {
			if err := f.client.ReadInConfig(); err != nil {
				return
			}
			// 遍历比对监听项值发生变化后回调
			for index, item := range f.watchList {
				newValue := fmt.Sprintf("%v", f.client.Get(item.key))
				if newValue != item.value {
					f.watchList[index].value = newValue
					item.cb(item.key, newValue)
				}
			}
		})
	})

	value, _ := f.Get(group, key)
	f.watchList = append(f.watchList, watchInfo{
		key:   f.Key(group, key),
		value: value,
		cb:    cb,
	})
	return nil
}

// --------------------------------

var (
	fileOnce sync.Once
	file     *fileInstance
)

// NewFile ...
func NewFile(confValue interface{}) *fileInstance {
	fileOnce.Do(func() {
		conf, ok := confValue.(*FileConfig)
		if !ok {
			panic("invalid file config")
		}
		_, err := os.Stat(conf.Path)
		if err != nil && !os.IsExist(err) {
			panic("config file not exist")
		}

		v := viper.New()
		v.AddConfigPath(path.Dir(conf.Path))
		v.SetConfigName(strings.Replace(path.Base(conf.Path), path.Ext(conf.Path), "", -1))
		v.SetConfigType(strings.Replace(path.Ext(conf.Path), ".", "", -1))
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}

		file = &fileInstance{
			conf:   conf,
			client: v,
		}
	})
	return file
}
