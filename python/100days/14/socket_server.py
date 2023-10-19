from socket import socket, SOCK_STREAM, AF_INET
from datetime import datetime


def main():
    server = socket(AF_INET, SOCK_STREAM)
    host = ('192.168.240.10', 6789)
    server.bind(host)
    # listen设定连接队列的大小
    server.listen(512)
    print('服务器开始监听...')
    while True:
        client, addr = server.accept()
        print(str(addr) + "连接到服务器.")
        # send data to client
        client.send(str(datetime.now()).encode('utf-8'))
        client.close()


if __name__ == "__main__":
    main()

