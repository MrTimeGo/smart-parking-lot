package s3

import (
	"context"
	"fmt"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/internal/s3/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

type CarStorage struct {
	s3     *minio.Client
	bucket string
}

func New(cfg config.S3Config) *CarStorage {
	client, err := minio.New(cfg.Credentials.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV2(cfg.Credentials.SecretKey, cfg.Credentials.AccessKey, ""),
	})
	if err != nil {
		panic(errors.Wrap(err, "failed to create s3 client"))
	}

	return &CarStorage{s3: client, bucket: cfg.Bucket}
}

func (c *CarStorage) Exists(car string) (bool, error) {
	_, err := c.s3.StatObject(context.Background(), c.bucket, car, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (c *CarStorage) ListCars() ([]string, error) {
	var cars []string
	doneCh := make(chan struct{})
	defer close(doneCh)

	for object := range c.s3.ListObjects(context.Background(), c.bucket, minio.ListObjectsOptions{Recursive: true}) {
		if object.Err != nil {
			return nil, object.Err
		}

		cars = append(cars, object.Key)
	}

	return cars, nil
}

func (c *CarStorage) CraftPath(car string) string {
	return fmt.Sprintf("%s/%s/%s", c.s3.EndpointURL().String(), c.bucket, car)
}
