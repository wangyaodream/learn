import asyncio
from websockets.sync.client import connect



def hello():
    with connect("ws://localhost:8765") as websocket:
        websocket.send("hello")
        message = websocket.recv()
        print(f"Received: {message}")


hello()


