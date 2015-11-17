package stm

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func NewS3Adapter() *S3Adapter {
	return &S3Adapter{ACL: "public-read"}
}

type S3Adapter struct {
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	Region             string
	Bucket             string
	ACL                string
}

func (adp *S3Adapter) Write(loc *Location, data []byte) {
	var reader io.Reader = bytes.NewReader(data)

	if GzipPtn.MatchString(loc.Filename()) {
		var writer *io.PipeWriter

		reader, writer = io.Pipe()
		go func() {
			gz := gzip.NewWriter(writer)
			io.Copy(gz, bytes.NewReader(data))

			gz.Close()
			writer.Close()
		}()
	}

	creds := credentials.NewEnvCredentials()
	creds.Get()

	sess := session.New(&aws.Config{
		Credentials: creds, Region: &adp.Region})

	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(adp.Bucket),
		Key:    aws.String(loc.PathInPublic()),
		ACL:    aws.String(adp.ACL),
		Body:   reader,
	})

	if err != nil {
		log.Fatal("[F] S3 Upload file Error:", err)
	}
}
