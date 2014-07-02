import socket, sys

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

s.connect(("127.0.0.1", 6226))
while True:
    print "Enter data to transmit:"
    data = sys.stdin.readline().strip()
    if len(data) > 0:
        s.sendall(data)
        buf = s.recv(2048)
        if not len(buf):
            break
        print "Server replies: ",
        sys.stdout.write(buf)
        print "\n"