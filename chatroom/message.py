'''
 header will receive 4 bytes and maxBytes can be stored as 512 bytes
 header will store the body length, which is the current body length
 data will be store with the header+bodyLength with maximum size of header+maxBytes

 client will attempt to send message ->
 it will encode header and put message into the data and send data
 to the server , get the message, decodes the header and get the bodyLength from the header
 and hence complete body then server send the message to all the clients connected to that room
    
'''

class Message:
    data: str
    body_length: int

    def __init__(self,message: str):
        self.body_length = len(message)
        self.encode_header()
        self.data = message

    def __repr__(self):
        print(f"Message received: {self.data}")
