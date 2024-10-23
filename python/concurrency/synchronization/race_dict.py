import asyncio

class MockSocket:
    def __init__(self):
        self.socket_closed = False
    
    async def send(self, msg: str):
        if self.socket_closed:
            raise Exception("Socket is closed!")
        print(f"Sending: {msg}")
        await asyncio.sleep(1)
        print(f'Sent: {msg}')

    def close(self):
        self.socket_closed = True



