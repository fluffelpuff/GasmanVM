package coreservice

import (
	"net"
)

type MemoryController struct {
	serverListener net.Listener
}

func (o *MemoryController) Close() {
	o.serverListener.Close()
}

func (o *MemoryController) AcceptConnection() (net.Conn, error) {
	conn, err := o.serverListener.Accept()
	if err != nil {
		return nil, err
	}
	return conn, err
}
