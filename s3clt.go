package s3clt

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
)

// amazon s3 - How to save data streams in S3? aws-sdk-go example not working? - Stack Overflow https://stackoverflow.com/questions/43595911/how-to-save-data-streams-in-s3-aws-sdk-go-example-not-working

type reader struct {
	r io.Reader
}

func (r *reader) Read(p []byte) (int, error) {
	return r.r.Read(p)
}

func getSession(args []string) (sess *session.Session, bucket string, key string) {
	var err error
	var region string
	if len(args) < 2 {
		panic("Too few arguments")
	}
	if len(args) >= 3 {
		region, args = args[0], args[1:]
	}
	bucket, args = args[0], args[1:]
	key, args = args[0], args[1:]
	sessOpt := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}
	if region != "" {
		sessOpt.Config.MergeIn(&aws.Config{
			Region: aws.String(region),
		})
	}
	sess, err = session.NewSessionWithOptions(sessOpt)
	if err != nil {
		panic(err)
	}
	if *(sess.Config.Region) == "" {
		data := ec2metadata.New(sess)
		metadataRegion, err := data.Region()
		if err == nil {
			sess.Config.MergeIn(&aws.Config{
				Region: aws.String(metadataRegion),
			})
		}
	}
	if *(sess.Config.Region) == "" {
		panic("Region missing")
	}
	return sess, bucket, key
}

type Command int8

const (
	CommandUnknown Command = iota
	CommandGet
	CommandPut
)

func Run(command Command, args []string) {
	var err error
	sess, bucket, key := getSession(args)
	switch command {
	case CommandGet:
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
	case CommandPut:
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
	default:
		panic("Unknown command")
	}
}
