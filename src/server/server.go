
package server

import (
    "net"
)

const ( maxRead = 1024 )


type IAuth interface {
    SignIn(username, password string)(bool)
}

func canSignIn(auth IAuth, username, password string) bool {
    return auth.SignIn(username, password)
}


type IConnectionManager interface {
    Accepted(msg string) error;
    Attach(conn net.Conn) error;
}

type IConnectionHandler interface {
    MessageReceived(conn net.Conn, msg []byte, length int) error;
}

type Server struct {
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
    if err == nil {
        listener, _ := net.ListenTCP("tcp", serverAddr)
        println("Listening to: ", listener.Addr().String())
        for {
            conn, err := listener.Accept()
            if err == nil {
                go s.acceptHandler(conn)
            } else {
                println(err)
            }
        }
    }
}

func (s Server) acceptHandler(conn net.Conn) {

    var buffer []byte = make([]byte, maxRead + 1)
    length, err := conn.Read(buffer[0 : maxRead])
    buffer[maxRead] = 0 // to prevent overflow

    if err != nil {
        msg := string(buffer[0: length])
        if s.connMgr.Accepted(msg) != nil {
            s.connMgr.Attach(conn)
        }
    }
}

func connectionHandler(conn net.Conn) {
    connFrom := conn.RemoteAddr().String()
    println("Connection from: ", connFrom)
    
    // TODO: Maybe for 
    for {
        var buffer []byte = make([]byte, maxRead + 1)
        length, err := conn.Read(buffer[0 : maxRead])
        buffer[maxRead] = 0 // to prevent overflow
        println("Received")
        switch err {
        case nil:
            handleMessage(buffer, length)
            conn.Write([]byte("a"))

        default:
            goto DISCONNECT
        }
    }

    DISCONNECT:
    err := conn.Close()
    if err != nil {

    }
    println("Closed connection:" , connFrom)
}

func handleMessage(msg []byte, length int) {
    println(string(msg[0: length]))
    // TODO: get the message type
    // TODO: if the message is chat message, parse the target
    // TODO: Get the target channel, send the reorgnized message.
}