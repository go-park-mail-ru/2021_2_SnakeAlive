package usecase

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/twinj/uuid"
	"snakealive/m/pkg/hasher"
)

type MediaUsecase interface {
	UploadFile(ctx context.Context, file io.ReadSeeker, ext string) (filename string, err error)
}

type mediaUsecase struct {
	gen           uuid.UUID
	hasher        hasher.Hasher
	defaultBucket string
	fileEndpoint  string
	client        *s3.S3
}

func (m *mediaUsecase) UploadFile(ctx context.Context, file io.ReadSeeker, ext string) (filename string, err error) {
	filename = m.hasher.EncodeString(m.gen.String())
	if _, err = m.client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Body:   file,
		Key:    aws.String(filename + ext),
		Bucket: aws.String(m.defaultBucket),
		ACL:    aws.String("public-read-write"),
	}); err != nil {
		return filename, err
	}

	return m.fileEndpoint + filename + ext, err
}

func NewMediaUsecase(
	gen uuid.UUID, client *s3.S3, hasher hasher.Hasher,
	defaultBucket, fileEndpoint string) MediaUsecase {
	return &mediaUsecase{
		gen:           gen,
		client:        client,
		defaultBucket: defaultBucket,
		fileEndpoint:  fileEndpoint,
		hasher:        hasher,
	}
}
