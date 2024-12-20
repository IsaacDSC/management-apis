package adapter

import (
	"context"
	"mime/multipart"
)

type PostmanCollectionRepository interface {
	SaveFile(ctx context.Context, collectionFilePath, envFilePath string, collectionFile, environmentFile multipart.File) error
	UnlinkFiles(ctx context.Context, files []string) error
}
