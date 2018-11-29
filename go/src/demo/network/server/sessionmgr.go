package server

import (
	"net"
	"sync"
)

type SessionMgr struct {
	sessionMap    sync.Map
	sessionNum    uint32
	maxSessionNum uint32
}

func (p *SessionMgr) GetNewSession(stream Streamer, conn net.Conn) Sessioner {
	session := NewSession(conn, stream, &sessionconfig{})
	p.sessionMap.Store(session.ID(), session)
	p.sessionNum++
	return session
}
