import json
import socket
socketstream=socket.socket(socket.AF_INET,socket.SOCK_STREAM)
server=("localhost",1234)
socketstream.connect(server)
op=['Add','Multiply','Subtract','Divide']
while True:
    try:
        print("Enter the Command in the given format [Add/Multiply/Subtract/Divide]<space>[num1]<space>num2")
        x=input()
        l=x.split(sep=" ")
        if (([(op.index(i)) for i in op if i==l[0]])==[]):
            print("Invalid Operation")
            continue
        request_param={
            'A':int(l[1]),
            'B':int(l[-1])
            }
        request={
            'method':f'ExportVariable.{l[0]}',
            'params':[request_param],            
            }
        construct_request=json.dumps(request)
        socketstream.sendall(construct_request.encode())
        response=json.loads(socketstream.recv(512).decode())
        print(response)
        print(f"{l[0]}:{response['result']}")
    except Exception :
        print(response['error'])
        continue
