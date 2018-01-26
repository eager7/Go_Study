package remote

import (
	"../storage"
	"fmt"
	"context"
	"github.com/pkg/errors"
	"github.com/eager7/go/mlog"
)

type Args struct {
	Key   interface{}
	Value interface{}
}

type Method struct {
	s storage.Storager
}

func NewMethod()(*Method){
	m := &Method{}
	m.s = storage.New()
	return m
}

func (m *Method)Add(ctx context.Context, key, value string)(reply int32, err error){
	mlog.Debug.Println("Add Method key:", key, "value:", value)
	err = m.s.Set(key, value)
	if err != nil {
		fmt.Println(err)
		reply = -1
		return
	}
	fmt.Println("Add key:" + key + "value:" + value)
	reply = 0
	return
}

func (m *Method)Get(ctx context.Context, key string)(value string, err error){
	mlog.Debug.Println("Get Method:", key)
	v, err := m.s.Get(key)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("can't find")
	}
	if _,ok := v.(string);!ok{
		fmt.Println("format error")
		return "", errors.New("format error")
	}
	fmt.Println("get key:" + key + "value:" + v.(string))

	return v.(string), nil
}