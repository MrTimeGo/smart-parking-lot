#!/usr/bin/env python
import pika


class RpcServer:
    def __init__(self, host, username, password):
        credentials = pika.PlainCredentials(username, password)
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(
                host=host,
                credentials=credentials,
        ))

        self.channel = self.connection.channel()

    def start_listening(self, queue_name, on_message):
        self.on_message = on_message
        self.channel.queue_declare(queue=queue_name)

        self.channel.basic_qos(prefetch_count=1)
        self.channel.basic_consume(queue='rpc_queue', on_message_callback=self.on_request)

        print(" [x] Awaiting RPC requests")
        self.channel.start_consuming()

    def on_request(self, ch, method, props, body):
        print(" [x] Received RPC request")
        response = self.on_message(body)

        ch.basic_publish(exchange='',
                         routing_key=props.reply_to,
                         properties=pika.BasicProperties(correlation_id= \
                                                             props.correlation_id),
                         body=str(response))
        ch.basic_ack(delivery_tag=method.delivery_tag)
        print(f" [x] Responded with {response}")