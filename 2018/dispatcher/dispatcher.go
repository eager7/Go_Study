// Copyright 2018 The go-ecoball Authors
// This file is part of the go-ecoball.
//
// The go-ecoball is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ecoball is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ecoball. If not, see <http://www.gnu.org/licenses/>.

package dispatcher

import (
	"fmt"
	"github.com/eager7/go_study/2018/dispatcher/pb"
	"gx/ipfs/QmdbxjQWogRCHRaxhhGnYdT1oQJzL9GdqSKzCdqWr85AP2/pubsub"
	"io"
)

const (
	bufferSize = 16
	errorStr   = "dispatcher is not ready"
)

var (
	dispatcher *Dispatcher
)

type Exportable interface {
	ToProtoV1() *pb.Message
	ToNetV1(w io.Writer) error
}

type NetMsg interface {
	Type() pb.MsgType
	Data() []byte
	//Exportable
}

func InitMsgDispatcher() {
	if dispatcher == nil {
		dispatcher = &Dispatcher{
			pubsub.New(bufferSize),
		}
	}
}

type Dispatcher struct {
	ps *pubsub.PubSub
}

func (ds *Dispatcher) publish(msg NetMsg) {
	ds.ps.Pub(msg, msg.Type().String())
}

func (ds *Dispatcher) subscribe(msg ...pb.MsgType) chan interface{} {
	var msgStr []string
	for _, msg := range msg {
		msgStr = append(msgStr, msg.String())
	}
	if len(msgStr) > 0 {
		return ds.ps.Sub(msgStr...)
	}

	return nil
}

func (ds *Dispatcher) unsubscribe(chn chan interface{}, msgType ...pb.MsgType) {
	var msgStr []string
	for _, msg := range msgType {
		msgStr = append(msgStr, msg.String())
	}

	ds.ps.Unsub(chn, msgStr...)
}

// Not safe to call more than once.
func (ds *Dispatcher) shutdown() {
	// shutdown the pub sub.
	ds.ps.Shutdown()
}

func Subscribe(msg ...pb.MsgType) (chan interface{}, error) {
	if dispatcher == nil {
		return nil, fmt.Errorf(errorStr)
	}
	return dispatcher.subscribe(msg...), nil
}

func UnSubscribe(chn chan interface{}, msg ...pb.MsgType) error {
	if dispatcher == nil {
		return fmt.Errorf(errorStr)
	}
	dispatcher.unsubscribe(chn, msg...)

	return nil
}

func Publish(msg NetMsg) error {
	if dispatcher == nil {
		return fmt.Errorf(errorStr)
	}
	dispatcher.publish(msg)

	return nil
}
