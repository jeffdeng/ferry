package main

import (
	"../xmpp"
	"errors"
	"redis"
	"net"
)



type BizConnection struct {
	
}


func (handler BizConnection) MessageReceived(conn net.Conn, msg []byte, length int) error {

	return nil;
}

func (handler BizConnection) SignIn(username, password string)(bool) {

    spec := redis.DefaultSpec().Host("127.0.0.1").Port(6380)
    client, _ := redis.NewSynchClientWithSpec(spec)

    value, _ := client.Llen("s:q")
    print(value)
    return true
}

type BizConnectionManager struct {
	connections map[string]net.Conn
}

func (connMgr BizConnectionManager) initialize() error {
	
	connMgr.connections = make(map[string]net.Conn)
	return nil
}

func (connMgr BizConnectionManager) Accepted(msg string, conn net.Conn) error {
	print(msg)
	// TODO: Check session
	return nil
}

func (connMgr BizConnectionManager) Attach(conn net.Conn) error {
	print(conn)
	return nil
}

func (connMgr BizConnectionManager) Lookup(name string) (conn net.Conn, err error) {
	conn, found := connMgr.connections[name]
	if found {
		return conn, nil
	}
	return nil, errors.New("")
}

func main() {

	server := new(xmpp.Server);
	server.Derived = server
	connMgr := BizConnectionManager{}
	connMgr.initialize()
	server.Start("127.0.0.1:6226", connMgr)
	if server != nil {

	}
}