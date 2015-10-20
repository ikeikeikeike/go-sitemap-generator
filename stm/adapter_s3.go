package stm

import "os"

func NewS3Adapter() *S3Adapter {
	adp := &S3Adapter{
		AwsAccessKeyId: os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
	return adp
}

type S3Adapter struct {
	AwsAccessKeyId     string
	AwsSecretAccessKey string
}

func (adp *S3Adapter) Write(loc *Location, data []byte) {}
