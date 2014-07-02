
package main

import (
    "net"
)

const ( maxRead = 1024)

func main() {

    println("A")
    initServer("127.0.0.1:6226")
}

func initServer(address string) {
    serverAddr, err := net.ResolveTCPAddr("tcp", address)
    if err == nil {
        listener, _ := net.ListenTCP("tcp", serverAddr)
        println("Listening to: ", listener.Addr().String())
        for {
            conn, err := listener.Accept()
            if err == nil {
                go connectionHandler(conn)
            } else {
                println(err)
            }
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