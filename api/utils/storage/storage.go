package storage

import (
	"context"
	"errors"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type IStorage interface {
	UplodaImage(ctx context.Context, imagePath string, publicID string) (*uploader.UploadResult, error)
}

func NewStorage(storageType string) (IStorage, error) {
	switch storageType {
	case "cloudinary":
		return newCloudinaryStorage()
	default:
		return nil, errors.New("invalid storage type")
	}
}
