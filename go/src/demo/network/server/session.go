package server

import (
	"net"
	"sync/atomic"
)

var sesssionID uint32

type Sessioner interface {
	ID() uint32
	Recv() (interface{}, error)
	Send(interface{}) error
	Close()
}

type sessionconfig struct {
}

type Session struct {
	*sessionconfig
	stream   Streamer
	id       uint32
	conn     net.Conn
	userDate interface{}
}

func NewSession(conn net.Conn, stream Streamer, cfg *sessionconfig) Sessioner {
	return &Session{
		id:            atomic.AddUint32(&sesssionID, 1),
		conn:          conn,
		stream:        stream,
		sessionconfig: cfg,
	}
}

func (p *Session) ID() uint32 {
	return p.id
}

func (p *Session) Send(data interface{}) error {
	return p.stream.Send(data)
}

func (p *Session) Recv() (interface{}, error) {
	return p.stream.Recv()
}

func (p *Session) Close() {
	p.stream.Close()
	p.conn = nil
}
