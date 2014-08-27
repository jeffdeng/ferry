import socket, sys

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

s.connect(("127.0.0.1", 6226))
s.sendall("<stream>")
buf = s.recv(2048)
sys.stdout.write(buf + "\n")
while True:
    print "Enter data to transmit:"
    data = sys.stdin.readline().strip()
    if len(data) > 0:
        if data == "bye":
            s.sendall("</stream>")
        else:
            s.sendall(data)
        buf = s.recv(2048)
        if not len(buf) or buf == '</stream>':
            break
        print "Server replies: ",
        sys.stdout.write(buf)
        print "\n"