package adapter

import (
	"bff/internal/management/dto/management"
	"bff/internal/management/dto/postman"
	"context"
	"mime/multipart"
)

type PostmanCollectionService interface {
	CreateFile(ctx context.Context, serviceName string, collectionFile, environmentFile multipart.File) (map[string]string, error)
	UnlinkFiles(ctx context.Context, files map[string]string) error
	GetInfos(ctx context.Context, input map[string]string) (*management.CollectionDto, *postman.Environment, error)
}
