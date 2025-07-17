import asyncio

clients = set()

SERVER_IP = "0.0.0.0"
SERVER_PORT = 8888

async def handle_client(reader, writer):
    addr = writer.get_extra_info('peername')
    print(f"New connection from {addr}")
    clients.add(writer)

    try:
        # send the data continuosuly
        while data := await reader.readline():
            message = data.decode().strip()
            print(f"{addr} : {message}")

            for client in clients:
                if client != writer:
                    client.write(f"{addr} : {message}".encode())
                    await client.drain()
    except asyncio.CancelledError:
        pass

    finally:
        print(f"Disconnected {addr}")
        clients.remove(writer)
        writer.close()
        await writer.wait_closed() 

async def main():
    
    server = await asyncio.start_server(handle_client,SERVER_IP,SERVER_PORT)
    addr = server.sockets[0].getsockname()
    print(f"Serving on {addr}")
    async with server:
        await server.serve_forever()



if __name__ == '__main__':
    asyncio.run(main())
