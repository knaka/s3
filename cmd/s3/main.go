package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
	"os"
	"path"
)

// amazon s3 - How to save data streams in S3? aws-sdk-go example not working? - Stack Overflow https://stackoverflow.com/questions/43595911/how-to-save-data-streams-in-s3-aws-sdk-go-example-not-working

type reader struct {
	r io.Reader
}

func (r *reader) Read(p []byte) (int, error) {
	return r.r.Read(p)
}

func getSession() (*session.Session, string, string, string) {
	var command, region, bucket, key string
	base := path.Base(os.Args[0])
	if base == "s3get" {
		command = "get"
	} else if base == "s3put" {
		command = "put"
	}
	args := os.Args[1:]
	if len(args) < 3 {
		panic("Too few arguments")
	}
	if command == "" {
		command, args = args[0], args[1:]
	}
	if len(args) >= 3 {
		region, args = args[0], args[1:]
	}
	bucket, args = args[0], args[1:]
	key, args = args[0], args[1:]

	log.Println("command:", command)
	log.Println("bucket:", bucket)
	log.Println("key:", key)

	sessOpt := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}
	if region != "" {
		sessOpt.Config.MergeIn(&aws.Config{
			Region: aws.String(region),
		})
	}
	sess, err := session.NewSessionWithOptions(sessOpt)
	if err != nil {
		panic(err)
	}
	if *(sess.Config.Region) == "" {
		data := ec2metadata.New(sess)
		mdRegion, err := data.Region()
		if err == nil {
			sess.Config.MergeIn(&aws.Config{
				Region: aws.String(mdRegion),
			})
		}
	}
	if *(sess.Config.Region) == "" {
		panic("Region missing")
	}
	log.Println("region:", *(sess.Config.Region))
	return sess, command, bucket, key
}

func main() {
	var err error
	sess, command, bucket, key := getSession()
	if command == "get" {
		client := s3.New(sess)
		result, err := client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
		if err != nil {
			fmt.Println("failed to get object,", err)
			return
		}
		defer (func() { _ = result.Body.Close() })()
		_, err = io.Copy(os.Stdout, result.Body)
		if err != nil {
			fmt.Println("failed to write to stdout,", err)
			return
		}
	}
	if command == "put" {
		uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {})
		_, err = uploader.UploadWithContext(context.Background(), &s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   &reader{os.Stdin},
		})
		if err != nil {
			log.Panicf("panic 02596e7 (%v)", err)
		}
		return
	}
}
