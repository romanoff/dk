package remote

import (
	"bytes"
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

func (self *S3) Push(filepath, destination string) error {
	bucket := self.getBucket()
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	perm := s3.BucketOwnerFull
	return bucket.Put(destination, content, "application/x-tar", perm)
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
