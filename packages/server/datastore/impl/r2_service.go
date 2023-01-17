package impl

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PureML-Inc/PureML/server/datastore/dbmodels"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2Secrets struct {
	AccountId       string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	PublicURL       string
}

func (r2secrets *R2Secrets) Load(secrets []dbmodels.Secret) error {
	for _, secret := range secrets {
		switch strings.ToUpper(secret.Name) {
		case "R2_ACCOUNT_ID":
			r2secrets.AccountId = secret.Value
		case "R2_ACCESS_KEY_ID":
			r2secrets.AccessKeyId = secret.Value
		case "R2_ACCESS_KEY_SECRET":
			r2secrets.AccessKeySecret = secret.Value
		case "R2_BUCKET_NAME":
			r2secrets.BucketName = secret.Value
		}
	}
	if r2secrets.AccountId == "" || r2secrets.AccessKeyId == "" || r2secrets.AccessKeySecret == "" || r2secrets.BucketName == "" {
		return fmt.Errorf("r2 secrets not found")
	}
	return nil
}

func (r2secrets *R2Secrets) CreateBucketIfNotExists() error {
	client := GetR2Client(*r2secrets)
	client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &r2secrets.BucketName,
	}) 
	// NEED TO CHECK IF RESPONSE IS 200 OR 409 THEN SUCCESS ELSE FAIL
	return nil
}

func GetR2Client(secrets R2Secrets) *s3.Client {
	var accountId = secrets.AccountId
	var accessKeyId = secrets.AccessKeyId
	var accessKeySecret = secrets.AccessKeySecret

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	return client
}
