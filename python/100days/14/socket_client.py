from socket import socket


def main():
    # 创建套接字对象
    client = socket()
    # 连接到服务器
    client.connect(('192.168.240.10', 6789))
    # 从服务器获取数据
    print(client.recv(1024).decode('utf-8'))
    client.close()


if __name__ == "__main__":
    main()

