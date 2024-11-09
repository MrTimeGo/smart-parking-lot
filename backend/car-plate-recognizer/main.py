from recognizer import recognize_image
from rpc_server import RpcServer
import os
from dotenv import load_dotenv

load_dotenv()

HOSTNAME = os.getenv("RABBITMQ_HOSTNAME")
USERNAME = os.getenv("RABBITMQ_USERNAME")
PASSWORD = os.getenv("RABBITMQ_PASSWORD")

QUEUE_NAME = os.getenv("RABBITMQ_QUEUE_NAME")




def main():
    print(USERNAME, PASSWORD, HOSTNAME, QUEUE_NAME)
    rpc_server = RpcServer(HOSTNAME, USERNAME, PASSWORD)
    rpc_server.start_listening(QUEUE_NAME, recognize_image)


if __name__ == "__main__":
    main()