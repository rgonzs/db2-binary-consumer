package services

import (
	"bytes"
	"context"
	"crypto/md5"
	"db2-binary-consumer/configuration"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type MessageAttachment struct {
	Binario     []byte
	TipoBinario string
}

type DocumentMessage struct {
	ID                  int
	SERIEDOCUMENTO      string
	SECUENCIALDOCUMENTO int
	TIPODOCUMENTO       string
	TIPOBINARIO         string
	IDEMISOR            string
	NOMBREARCHIVO       string
	FECHAEMISION        string
	BINARIO             []byte
	Attachments         []MessageAttachment
}

func UploadToStorage(pathFile string, bodyFile []byte) (string, error) {
	client := configuration.S3Connection()
	md5 := getContentMd5(bodyFile)
	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:     aws.String("pse-db2-storage"),
		Key:        aws.String(pathFile),
		Body:       bytes.NewReader(bodyFile),
		ContentMD5: &md5,
	})
	if err != nil {
		return "", err
	}
	return md5, nil
}

func getContentMd5(body []byte) string {
	encoded := fmt.Sprintf("%x", md5.Sum(body))
	sumHex, _ := hex.DecodeString(encoded)
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(sumHex)))
	base64.StdEncoding.Encode(eb, sumHex)
	return string(eb)
}
