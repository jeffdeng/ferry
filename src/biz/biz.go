
package main

import (
	"../server"
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

}

func (connMgr BizConnectionManager) Accepted(msg string) error {
	print(msg)
	return nil
}

func (connMgr BizConnectionManager) Attach(conn net.Conn) error {
	
	return nil
}


func main() {

	server := new(server.Server);
	connMgr := BizConnectionManager{}
	server.Start("127.0.0.1:6226", connMgr)
}