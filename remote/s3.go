package remote

import (
	"bytes"
	a_aws "github.com/aws/aws-sdk-go/aws"
	a_session "github.com/aws/aws-sdk-go/aws/session"
	a_s3 "github.com/aws/aws-sdk-go/service/s3"
	"io"
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"os"
)

type S3 struct {
	Bucket    string
	AccessKey string
	SecretKey string
	Region    string
}

func (self *S3) Setup(config *Config) {
	self.Bucket = config.Bucket
	self.AccessKey = config.Access_Key
	self.SecretKey = config.Secret_Key
	self.Region = config.Region
}

func (self *S3) getBucket() *s3.Bucket {
	auth := aws.Auth{
		AccessKey: self.AccessKey,
		SecretKey: self.SecretKey,
	}
	region := aws.Regions[self.Region]
	connection := s3.New(auth, region)
	return connection.Bucket(self.Bucket)
}

func (self *S3) FilesList() ([]string, error) {
	bucket := self.getBucket()
	res, err := bucket.List("", "", "", 1000)
	if err != nil {
		return nil, err
	}
	files := []string{}
	for _, v := range res.Contents {
		files = append(files, v.Key)
	}
	return files, nil
}

func (self *S3) GetS3Session() (*a_s3.S3, error) {
	err := os.Setenv("AWS_ACCESS_KEY_ID", self.AccessKey)
	if err != nil {
		return nil, err
	}
	err = os.Setenv("AWS_SECRET_ACCESS_KEY", self.SecretKey)
	if err != nil {
		return nil, err
	}
	return a_s3.New(a_session.New(),
		&a_aws.Config{
			Region: a_aws.String(self.Region),
		},
	), nil
}

func (self *S3) Push(filepath, destination string) error {
	session, err := self.GetS3Session()
	if err != nil {
		return err
	}
	content, err := ioutil.ReadFile(filepath)
	params := &a_s3.PutObjectInput{
		Bucket:               a_aws.String(self.Bucket), // Required
		Key:                  a_aws.String(destination), // Required
		ACL:                  a_aws.String("bucket-owner-read"),
		Body:                 bytes.NewReader(content),
		ContentType:          a_aws.String("application/x-tar"),
		ServerSideEncryption: a_aws.String("AES256"),
	}
	_, err = session.PutObject(params)
	if err != nil {
		return err
	}
	return nil
}

func (self *S3) Pull(filepath, destination string) error {
	bucket := self.getBucket()
	content, err := bucket.Get(filepath)
	if err != nil {
		return err
	}
	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	if _, err := io.Copy(destinationFile, bytes.NewReader(content)); err != nil {
		destinationFile.Close()
		return err
	}
	return destinationFile.Close()
}

func (self *S3) Remove(filepath string) error {
	bucket := self.getBucket()
	return bucket.Del(filepath)
}
