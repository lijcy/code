package server

import (
	"net"
)

//抽象数据接受发送模块
type Streamer interface {
	Recv() (interface{}, error)
	Send(interface{}) error
	Close()
}

//抽象数据发送对象的管理器
type StreamFactory interface {
	NewStream(net.Conn) Streamer
}

//包头
type HeadFormater interface {
	HeadLen() uint32
	Encode([]byte, uint32)
	Decode([]byte) uint32
}

//包头
type BodyFormater interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}
