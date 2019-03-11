package jsonrpc

import (
	"errors"
	"reflect"
	"sync"
	"unicode"
	"unicode/utf8"
)

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
}

type service struct {
	name     string
	receiver reflect.Value
	typ      reflect.Type
	methods  map[string]*methodType
}

type Server struct {
	serviceList sync.Map
}

func (s *Server) Register(receiver interface{}) error {
	_service := service{
		name:     "",
		receiver: reflect.ValueOf(receiver),
		typ:      reflect.TypeOf(receiver),
		methods:  nil,
	}
	sName := reflect.Indirect(reflect.ValueOf(receiver)).Type().Name()
	if sName == "" {
		return errors.New("rpc.Register: no service name for type " + _service.typ.String())
	}
	if !isExported(sName) {
		return errors.New("rpc.Register: type " + sName + " is not exported")
	}
	_service.name = sName
	//_service.methods[]
	return nil
}

// type isExported
func isExported(name string) bool {
	r, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(r)
}
