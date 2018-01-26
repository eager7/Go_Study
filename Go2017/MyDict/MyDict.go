package MyDict

import (
	"errors"
	"sync"
)

type Dict interface{
	Add(key interface{}, value interface{})
	Delete(key interface{})
	Update(key interface{}, value interface{})error
	Query(key interface{})(interface{},error)
}

type dict struct{
	mutex sync.RWMutex
	data map[interface{}]interface{}
}

func New()Dict{
	data := &dict{}
	data.data = make(map[interface{}]interface{})
	return data
}

func (d *dict)Add(key interface{}, value interface{}){
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.data[key] = value
}

func (d *dict)Delete(key interface{}){
	d.mutex.Lock()
	defer d.mutex.Unlock()
	delete(d.data, key)
}

func (d *dict)Update(key interface{}, value interface{})error{
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if _,ok := d.data[key];ok{
		d.data[key] = value
		return nil
	}
	return errors.New("not found")
}

func (d *dict)Query(key interface{})(interface{}, error){
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	if _,ok := d.data[key];ok{
		return d.data[key], nil
	}
	return nil, errors.New("not found")
}
