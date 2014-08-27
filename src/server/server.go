
package server

import (
    "net"
    "crypto/tls"
)

const ( maxRead = 1024 )


type IServer interface {
    MessageDispach(conn net.Conn) bool
}

type IConnectionManager interface {
   
    Accepted(msg string, conn net.Conn) error;
    Lookup(name string) (conn net.Conn, err error);
    Attach(conn net.Conn) error;
}



type IConnectionHandler interface {
    MessageReceived(conn net.Conn, msg []byte, length int) error;
}

type Server struct {
    Derived IServer
    address string
    
    connMgr IConnectionManager
}

func (s Server) Start(address string, connMgr IConnectionManager) error {
    s.address = address
    s.connMgr = connMgr
    s.initServer()
    return nil
}


func (s Server) initServer() {
    serverAddr, err := net.ResolveTCPAddr("tcp", s.address)
    if serverAddr == nil {
        
    }
    if err == nil {


        cert, err := tls.LoadX509KeyPair("e:/cacert.pem", "e:/privkey.pem")
        var config tls.Config
        if err == nil {
            config = tls.Config {Certificates: []tls.Certificate {cert}}
        }


        listener, _ := tls.Listen("tcp", s.address, &config)
        println("Server@" + listener.Addr().String())
        for {
            conn, err := listener.Accept()
            if err == nil {
                go s.Derived.MessageDispach(conn)
            } else {
                println(err)
            }
        }
    }
}

func (s Server) ReadMsg(conn net.Conn, maxRead int) (msg string, err error) {
    var buffer []byte = make([]byte, maxRead + 1)
    length, err := conn.Read(buffer[0 : maxRead])
    buffer[maxRead] = 0 // to prevent overflow

    if err == nil {
        msg := string(buffer[0: length])
        return msg, nil
    } else {
        return "", err
    }
}

func (s Server) MessageRoutine(conn net.Conn) {

    for {
        if s.Derived.MessageDispach(conn) != true {
            break
        }
    }
}
