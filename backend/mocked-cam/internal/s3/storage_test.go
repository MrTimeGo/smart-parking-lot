package s3

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Exists(t *testing.T) {
	client, err := minio.New("localhost:9000", &minio.Options{
		Creds: credentials.NewStaticV2("minioadmin", "minioadmin", ""),
	})
	if err != nil {
		t.Fatal(err)
	}

	storage := New(client, "cars")
	exists, err := storage.Exists("car1.jpg")
	if err != nil {
		t.Fatal(err)
	}

	assert.False(t, exists)

	exists, err = storage.Exists("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, exists)

	cars, err := storage.ListCars()
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, cars, 1)
	assert.Equal(t, "test.txt", cars[0])

}
