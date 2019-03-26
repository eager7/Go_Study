package viper

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Initialize(file string) error {
	viper.SetConfigFile(file)
	viper.SetDefault("key1", "value1")
	viper.SetDefault("key2", "value2")
	fmt.Println("config file:", viper.ConfigFileUsed())
	if _, err := os.Stat(viper.ConfigFileUsed()); os.IsNotExist(err) {
		if _, err := os.Create(viper.ConfigFileUsed()); err != nil {
			fmt.Println("create file err:", err)
			return err
		}
		if err := viper.WriteConfig(); err != nil {
			return err
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("read config file error:", err)
		return err
	}
	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
		})
	}()
	return nil
}
