package stm

import (
	"compress/zlib"
	"log"
	"os"
	"regexp"
)

var gzipPtn = regexp.MustCompile(".gz$")

func NewS3Adapter() *S3Adapter {
	adapter := &S3Adapter{}
	return adapter
}

type S3Adapter struct{
	AwsAccessKeyId = opts[:aws_access_key_id] || ENV['AWS_ACCESS_KEY_ID']
	AwsSecretAccess_key = opts[:aws_secret_access_key] || ENV['AWS_SECRET_ACCESS_KEY']
}

func (a *S3Adapter) Write(loc *Location, data []byte) {

}
