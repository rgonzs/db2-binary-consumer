package repository

import (
	"context"
	"db2-binary-consumer/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Attachment struct {
	Path       string
	ContentMd5 string
}

type Attachments struct {
	XMLDATA Attachment
	UBL     Attachment
	CDR     Attachment
}

type DocumentsRepository struct {
	Db2id               int    `bson:"db2id"`
	SerieDocumento      string `bson:"serieDocumento"`
	SecuencialDocumento int    `bson:"secuencialDocumento"`
	TipoDocumento       string `bson:"tipoDocumento"`
	FechaEmision        string `bson:"fechaEmision"`
	IdEmisor            string `bson:"idEmisor"`
	Attachments         Attachments
}

type DocumentResult struct {
	M_ID primitive.ObjectID `bson:"_id,omitempty"`
}

func (dr DocumentsRepository) Insert(ctx context.Context, collection *mongo.Collection) {
	date, _ := utils.ParseDateToLocal(dr.FechaEmision)
	doc := bson.D{
		{Key: "db2id", Value: dr.Db2id},
		{Key: "serieDocumento", Value: dr.SerieDocumento},
		{Key: "secuencialDocumento", Value: dr.SecuencialDocumento},
		{Key: "tipoDocumento", Value: dr.TipoDocumento},
		{Key: "idEmisor", Value: dr.IdEmisor},
		{Key: "fechaEmision", Value: date.Unix()},
		{Key: "attachments", Value: bson.D{
			{Key: "XMLDATA", Value: dr.Attachments.XMLDATA},
			{Key: "UBL", Value: dr.Attachments.UBL},
			{Key: "CDR", Value: dr.Attachments.CDR},
		}},
	}
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Println("Documento duplicado con db2 id: ", dr.Db2id)
		log.Println(err)
		return
	}
	log.Println("Documento procesado ", result.InsertedID)
}

func GetDocument(ctx context.Context, collection *mongo.Collection, db2id int) (string, bool) {
	opts := options.FindOne().SetProjection(bson.D{{Key: "db2id", Value: 1}, {Key: "_id", Value: 1}})
	doc := bson.D{{Key: "db2id", Value: db2id}}
	var result DocumentResult
	collection.FindOne(ctx, doc, opts).Decode(&result)
	return result.M_ID.Hex(), result.M_ID.IsZero()
}
