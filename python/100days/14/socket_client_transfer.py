from socket import socket
from json import loads
from base64 import b64decode


def main():
    client = socket()
    client.connect(('192.168.240.10', 5678))
    # 定义一个保存二进制数据的对象
    in_data = bytes()
    # 每次都接受1024字节
    data = client.recv(1024)
    while data:
        in_data += data
        data = client.recv(1024)

    my_dict = loads(in_data.decode('utf-8'))
    filename = my_dict['filename']
    filedata = my_dict['filedata']

    with open(f'result/{filename}', 'wb') as fp:
        fp.write(b64decode(filedata))

    print('图片已保存.')


if __name__ == "__main__":
    main()

