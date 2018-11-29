package server

import (
	"encoding/binary"
	"net"
)

//包头解析器
type testHeadFormat struct {
	headSize uint32
}


//草泥马

func (p *testHeadFormat) HeadLen() uint32 {
	return p.headSize
}

func (p *testHeadFormat) Encode(data []byte, size uint32) {
	binary.BigEndian.PutUint32(data, uint32(size)+4)
}

func (p *testHeadFormat) Decode(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

//包体解析器
type testBodyFormat struct {
}

func (p *testBodyFormat) Encode(data interface{}) ([]byte, error) {
	msg := data.(*[]byte)
	return *msg, nil
}

func (p *testBodyFormat) Decode(data []byte) (interface{}, error) {
	msg := &data
	return msg, nil
}

type testStreamFactory struct {
}

func NewTestStreamFactory() StreamFactory {
	return &testStreamFactory{}
}

func (p *testStreamFactory) NewStream(conn net.Conn) Streamer {
	return NewTestSteam(conn, &testHeadFormat{}, &testBodyFormat{})
}

type testStream struct {
	headFormat HeadFormater
	bodyFormat BodyFormater
	conn       net.Conn
}

func NewTestSteam(conn net.Conn, h HeadFormater, b BodyFormater) Streamer {
	return &testStream{
		conn:       conn,
		headFormat: h,
		bodyFormat: b,
	}
}

func (p *testStream) Send(data interface{}) error {
	msg, err := p.bodyFormat.Encode(data)
	if err != nil {
		return err
	}
	p.conn.Write(msg)
	return nil
}

func (p *testStream) Recv() (interface{}, error) {
	var buf []byte
	p.conn.Read(buf)
	msg, _ := p.bodyFormat.Decode(buf)
	return msg, nil
}

func (p *testStream) Close() {
	p.conn.Close()
}
