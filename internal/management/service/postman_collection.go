package service

import (
	"bff/internal/management/adapter"
	"bff/internal/management/dto/management"
	"bff/internal/management/dto/postman"
	"context"
	"fmt"
	"mime/multipart"
)

type PostmanCollection struct {
	repository adapter.PostmanCollectionRepository
}

func NewPostmanCollection(repository adapter.PostmanCollectionRepository) *PostmanCollection {
	return &PostmanCollection{repository: repository}
}

const (
	environment = "environment"
	collection  = "collection"
)

func (pc PostmanCollection) CreateFile(ctx context.Context, serviceName string, collectionFile, environmentFile multipart.File) (map[string]string, error) {
	collectionFilePath := fmt.Sprintf("tmp/%s_%s.json", serviceName, collection)
	envFilePath := fmt.Sprintf("tmp/%s_%s.json", serviceName, environment)

	if err := pc.repository.SaveFile(ctx, collectionFilePath, envFilePath, collectionFile, environmentFile); err != nil {
		return nil, err
	}

	return map[string]string{
		"collection":  collectionFilePath,
		"environment": envFilePath,
	}, nil
}

func (pc PostmanCollection) UnlinkFiles(ctx context.Context, files map[string]string) error {
	listFiles := make([]string, len(files))

	var i int
	for _, file := range files {
		listFiles[i] = file
		i++
	}

	return pc.repository.UnlinkFiles(ctx, listFiles)
}

func (pc PostmanCollection) GetInfos(ctx context.Context, input map[string]string) (*management.CollectionDto, *postman.Environment, error) {
	collectionDto, err := management.NewPostmanCollection(input["collection"])
	if err != nil {
		return nil, nil, fmt.Errorf("invalid postman collection: %v", err)
	}

	environmentDto, err := postman.NewEnvironment(input["environment"])
	if err != nil {
		return nil, nil, fmt.Errorf("invalid postman environment: %v", err)
	}

	return &collectionDto, &environmentDto, nil

}
