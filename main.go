package main

import (
	"db2-binary-consumer/configuration"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Cron("*/1 * * * *").Do(
		func() {
			log.Println("Starting to process documents")
		},
	)
	// scheduler.StartBlocking()

	forever := make(chan bool)

	mongoConfig := configuration.MongoClient{Uri: "mongodb+srv://admin:Yf2C9YdnNQBmZnan@cluster0.aftq1.mongodb.net/?retryWrites=true&w=majority"}
	client, ctx, cancel := mongoConfig.Connect()
	coll := client.Database("documents").Collection("history")

	defer mongoConfig.Close(client, ctx, cancel)

	// CPU THREADS
	for i := 1; i <= 10; i++ {
		connection, channel := configuration.RabbitConnect()
		// notify := connection.NotifyClose(make(chan *amqp.Error))

		defer connection.Close()
		defer channel.Close()

		msgs, err := channel.Consume("pse-documents-queue", "", false, false, false, false, nil)

		if err != nil {
			log.Println(err)
			log.Fatal("No se pudo conectar a la cola")
		}

		if err != nil {
			log.Fatal(err)
		}

		go ProcessDocument(ctx, coll, msgs)

	}

	// wg.Wait()
	<-forever
}
