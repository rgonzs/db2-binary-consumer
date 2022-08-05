package configuration

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func RabbitConnect() (*amqp.Connection, *amqp.Channel) {
	url := "amqps://hnocdams:q7ZQN6J2ZqDoR8615t2JHgpGc7w9Pktb@shark.rmq.cloudamqp.com/hnocdams"
	rabbitConnection, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to rabbitmq")

	channel, err := rabbitConnection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	channel.Qos(500, 0, false)
	log.Println("Rabbit Channel created")
	if err != nil {
		log.Fatal(err)
	}

	return rabbitConnection, channel
}
