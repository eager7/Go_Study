package storage

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/eager7/go/mlog"
)

type Storager interface {
	Set(key, value interface{})error
	Get(key interface{})(value interface{}, err error)
}

type storage struct{
	r redis.Conn
}

func New()Storager{
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil{
		fmt.Println(err)
		return nil
	}

	return &storage{c}
}

func (s *storage)Set(key, value interface{})error{
	_, err := s.r.Do("SET", key, value)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *storage)Get(key interface{})(value interface{}, err error){
	v, err := redis.String(s.r.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return v, nil
}
