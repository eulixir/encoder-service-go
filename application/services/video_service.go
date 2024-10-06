package services

import (
	"context"
	"encoder/application/repositories"
	"encoder/domain"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

// func NewVideoService() VideoService {

// }

func (v *VideoService) Download(bucketName string) error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)

	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(v.Video.FilePath)

	r, err := obj.NewReader(ctx)

	defer r.Close()

	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	var localStoragePath string = os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4"

	f, err := os.Create(localStoragePath)

	if err != nil {
		return err
	}

	_, err = f.Write(body)

	if err != nil {
		return err
	}

	defer r.Close()

	log.Print("video %w has been stored", v.Video.ID)

	return nil
}
