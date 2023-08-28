from threading import Thread
import socket


def echo(client: socket):
    while True:
        data = client.recv(1024)
        print(f'Received {data}, sending')
        client.sendall(data)


