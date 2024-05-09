import grpc
from concurrent import futures
import time
import script_server_pb2_grpc as pb2_grpc
import script_server_pb2 as pb2
from custom import custom_init, custom_handle

class ScriptServer(pb2_grpc.ScriptServer):

    def __init__(self, *args, **kwargs):
        pass

    def RunScript(self, request, context):

        message = request.req
        print(message)
        result = {'rsp': custom_handle(message)}

        return pb2.ScriptResponse(**result)



def serve():
    custom_init()
    print("serving started")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_ScriptServerServicer_to_server(ScriptServer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()