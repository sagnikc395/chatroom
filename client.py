# an minimal implementation of a client

import asyncio

CLIENT_IP = "127.0.0.1"
CLIENT_PORT = 8888

async def listen(reader):
    while data := await reader.readline():
        print(data.decode().strip())

async def main():
    print('chatroom client')

    [reader,writer] = await asyncio.open_connection(CLIENT_IP,CLIENT_PORT)

    asyncio.create_task(listen(reader))

    while True:
        msg = input()
        writer.write(f"{msg}\n".encode())
        await writer.drain()
    


if __name__ == '__main__':
    asyncio.run(main())
