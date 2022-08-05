package main

import (
	"context"
	"db2-binary-consumer/repository"
	"db2-binary-consumer/services"
	"db2-binary-consumer/utils"
	"encoding/json"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProcessDocument(ctx context.Context, collection *mongo.Collection, msgs <-chan amqp091.Delivery) {
	//const layout string = "2006-01-02T15:04:05Z07:00" // RFC3339

	for msg := range msgs {
		var document services.DocumentMessage
		var attachments repository.Attachments
		var path string
		json.Unmarshal(msg.Body, &document)
		dt, err := utils.ParseDateToLocal(document.FECHAEMISION)
		if err != nil {
			log.Println("Error al parsear la fecha:", err)
			msg.Ack(true)
			continue
		}

		mongoID, isMigrated := repository.GetDocument(context.TODO(), collection, document.ID)
		if !isMigrated {
			log.Println("Anteriormente migrado, id en mongodb", mongoID)
			msg.Ack(true)
			continue
		}

		for _, attachment := range document.Attachments {
			switch attachment.TipoBinario {
			case "UBL":
				path = fmt.Sprintf("/%d/UBL/%s/%s.zip", dt.Year(), document.IDEMISOR, document.NOMBREARCHIVO)
			case "XMLDATA":
				path = fmt.Sprintf("/%d/XMLDATA/%s/XMLDATA-%s.xml", dt.Year(), document.IDEMISOR, document.NOMBREARCHIVO)
			case "CDR":
				path = fmt.Sprintf("/%d/CDR/%s/R-%s.zip", dt.Year(), document.IDEMISOR, document.NOMBREARCHIVO)
			}
			out, err := services.UploadToStorage(path, attachment.Binario)
			if err != nil {
				log.Println("Error al subir al s3")
				log.Println(err)
				continue
			}
			att := repository.Attachment{
				Path:       path,
				ContentMd5: out,
			}
			switch attachment.TipoBinario {
			case "UBL":
				attachments.UBL = att
			case "XMLDATA":
				attachments.XMLDATA = att
			case "CDR":
				attachments.CDR = att
			}

		}

		documento := repository.DocumentsRepository{
			Db2id:               document.ID,
			SerieDocumento:      document.SERIEDOCUMENTO,
			SecuencialDocumento: document.SECUENCIALDOCUMENTO,
			TipoDocumento:       document.TIPODOCUMENTO,
			FechaEmision:        document.FECHAEMISION,
			IdEmisor:            document.IDEMISOR,
			Attachments:         attachments,
		}
		documento.Insert(context.TODO(), collection)
		log.Printf("Message processed: %s", document.NOMBREARCHIVO)
		msg.Ack(true)
	}
}
