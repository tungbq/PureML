package impl

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PureML-Inc/PureML/server/datastore/dbmodels"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Secrets struct {
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	PublicURL       string
}

func (s3secrets *S3Secrets) Load(secrets []dbmodels.Secret) error {
	for _, secret := range secrets {
		switch strings.ToUpper(secret.Name) {
		case "S3_ACCESS_KEY_ID":
			s3secrets.AccessKeyId = secret.Value
		case "S3_ACCESS_KEY_SECRET":
			s3secrets.AccessKeySecret = secret.Value
		case "S3_BUCKET_NAME":
			s3secrets.BucketName = secret.Value
		}
	}
	if s3secrets.AccessKeyId == "" || s3secrets.AccessKeySecret == "" || s3secrets.BucketName == "" {
		return fmt.Errorf("r2 secrets not found")
	}
	return nil
}

func (s3secrets *S3Secrets) CreateBucketIfNotExists() error {
	client := GetS3Client(*s3secrets)
	client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &s3secrets.BucketName,
	})
	// NEED TO CHECK IF RESPONSE IS 200 OR 409 THEN SUCCESS ELSE FAIL
	return nil
}

func GetS3Client(secrets S3Secrets) *s3.Client {
	var accessKeyId = secrets.AccessKeyId
	var accessKeySecret = secrets.AccessKeySecret

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	return client
}
