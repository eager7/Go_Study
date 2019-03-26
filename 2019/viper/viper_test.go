package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestViper(t *testing.T) {
	if err := Initialize("./config.toml"); err != nil {
		t.Fatal(err)
	}

	fmt.Println(viper.Get("key1"))
	fmt.Println(viper.Get("key2"))

	//make(chan interface{}, 0) <- struct{}{}
}
