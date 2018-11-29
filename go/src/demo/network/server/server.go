package server

import (
	"fmt"
	"net"
)

type Server interface {
	Close()
	DoEvent()
	SetUserData(interface{})
}

type SessionEventHandler interface {
	OnOpen(Sessioner, bool)
	OnClose(Sessioner)
	OnMessage(Sessioner, interface{})
}

type serverconfig struct {
	network string
	addr    string
	Factory StreamFactory
	Handler SessionEventHandler
}

type server struct {
	*SessionMgr
	*serverconfig
	l        net.Listener
	userData interface{}
}

func (p *server) Close() {
}

func (p *server) DoEvent() {
}

func (p *server) SetUserData(data interface{}) {
	p.userData = data
}

func NewTcpServer(cfg *serverconfig) Server {
	l, err := net.Listen(cfg.network, cfg.addr)
	if err != nil {
		fmt.Println("lister error , addr :", cfg.addr)
	}

	svr := &server{
		serverconfig: cfg,
		l:            l,
		SessionMgr:   &SessionMgr{maxSessionNum: 10000},
	}
	go svr.server()
	return svr
}

func (p *server) server() {
	for {
		conn, err := p.l.Accept()
		if err != nil {
			continue
		}
		stream := p.Factory.NewStream(conn)
		session := p.SessionMgr.GetNewSession(stream, conn)
		p.postOpen(session)
		go p.serverOne(session.(*Session))
	}
}

func (p *server) serverOne(session *Session) {
	for {
		msg, err := session.Recv()
		if err != nil {
			session.Close()
			return
		}
		p.postMsg(session, msg)
	}
}

func (p *server) postOpen(s Sessioner) {
	p.Handler.OnOpen(s, true)
}

func (p *server) postMsg(s Sessioner, msg interface{}) {
	p.Handler.OnMessage(s, msg)
}
