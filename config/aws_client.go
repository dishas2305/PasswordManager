package config

import (
	"fmt"
	"os"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sns"
	logger "github.com/sirupsen/logrus"
)

func GetSESClient() (*ses.SES, error) {
	logger.Info("Get client for aws-ses email")

	// Create a new session and specify an AWS Region.
	// Replace us-west-2 with the AWS Region you're using for Amazon SES.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_EMAIL_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_EMAIL_ACCESS_KEY_ID"), os.Getenv("AWS_EMAIL_SECRET_ACCESS_KEY"), ""),
	})

	fmt.Println("sess--->>", sess)
	if err != nil {
		logger.Error("error in ses NewSession ", err.Error())
		return nil, err
	}

	// Create an SES client in the session.
	svc := ses.New(sess)

	fmt.Println("svc---->>", svc)

	return svc, nil
}

func GetSNSClient() (*sns.SNS, error) {
	logger.Info("Get client for aws-sns")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_SNS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_SNS_ACCESS_KEY_ID"), os.Getenv("AWS_SNS_SECRET_ACCESS_KEY"), ""),
	})

	if err != nil {
		logger.Error("error in sns NewSession ", err.Error())
		return nil, err
	}

	// Create an SNS client in the session.
	svc := sns.New(sess)

	return svc, nil
}

func GetS3Client() (*s3manager.Uploader, error) {
	logger.Info("Get client for aws-s3")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_S3_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_S3_ACCESS_KEY_ID"), os.Getenv("AWS_S3_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		logger.Error("error in S3 NewSession ", err.Error())
		return nil, err
	}

	uploader := s3manager.NewUploader(sess)

	return uploader, nil
}

func GetS3ClientForDeleteObject() (*s3.S3, error) {
	logger.Info("Get client for aws-s3 delete object")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_S3_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_S3_ACCESS_KEY_ID"), os.Getenv("AWS_S3_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		logger.Error("error in S3 NewSession ", err.Error())
		return nil, err
	}

	// Create S3 service client
	svc := s3.New(sess)

	return svc, nil
}

func GetS3ClientForDownloadingFile() (*s3manager.Downloader, error) {
	logger.Info("Get client for aws-s3 downloadObject")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_S3_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_S3_ACCESS_KEY_ID"), os.Getenv("AWS_S3_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		logger.Error("error in  GetS3ClientForDownloadingFile S3 NewSession ", err.Error())
		return nil, err
	}

	downloader := s3manager.NewDownloader(sess)

	return downloader, nil
}
